package discovery

import "C"
import (
	"encoding/base64"
	"encoding/json"

	"github.com/limanmys/inventory-server/app/entities"
	"github.com/limanmys/inventory-server/internal/database"
	"github.com/limanmys/inventory-server/pkg/aes"
	"github.com/rainycape/dl"
)

func Start(discovery entities.Discovery) {
	// Open c-shared library
	lib, err := dl.Open("./wmi/wmi.so", 0)
	if err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when opening lib, err: "+err.Error())
		return
	}

	// Close library after function returned
	defer lib.Close()

	// Set, get function
	var get func(*C.char) *C.char

	// Open symbol
	err = lib.Sym("get", &get)
	if err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when getting symbol, err: "+err.Error())
		return
	}

	// Update discovery status
	discovery.UpdateStatus(entities.DiscoveryStatusInProgress, "Discovery in progress.")

	// Get profile
	var profile entities.Profile
	if err := database.Connection().Model(&profile).Where("id = ?", discovery.ProfileID).First(&profile).Error; err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when getting profile, err: "+err.Error())
		return
	}

	// Decrypt profile
	profile, err = aes.DecryptProfile(profile)
	if err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when decrypting profile, err: "+err.Error())
		return
	}

	// Run discovery function
	result, err := json.Marshal(entities.Arguments{
		IPRange:  discovery.IPRange,
		Username: profile.Username,
		Password: profile.Password,
	})
	if err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when running symbol, err: "+err.Error())
		return
	}

	// Unmarshal discovery result
	var discoveryResult entities.Result
	if err := json.Unmarshal([]byte(C.GoString(get(C.CString(string(result))))), &discoveryResult); err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when unmarshalling result, err: "+err.Error())
		return
	}

	// Decode b64 data
	data, err := base64.StdEncoding.DecodeString(discoveryResult.Output.(string))
	if err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when decoding result, err: "+err.Error())
		return
	}

	// Unmarshal assets
	var assets []entities.Asset
	if err := json.Unmarshal(data, &assets); err != nil {
		discovery.UpdateStatus(entities.DiscoveryStatusError, "error when unmarshalling assets, err: "+err.Error())
		return
	}

	var count int64
	// Create assets
	for _, asset := range assets {
		// Set discovery id
		asset.DiscoveryID = discovery.ID

		// Check is asset exists
		database.Connection().
			Model(&entities.Asset{}).Where("address = ?", asset.Address).Count(&count)
		if count > 0 {
			// If asset exists, update
			database.Connection().
				Model(&entities.Asset{}).Where("address = ?", asset.Address).Updates(&asset)
		} else {
			// Create asset
			database.Connection().Create(&asset)
		}

		// Create packages
		var assetPackages []*entities.Package
		for _, pkg := range asset.Packages {
			database.Connection().
				Model(&entities.Package{}).
				Where("name = ? and version = ?", pkg.Name, pkg.Version).Count(&count)

			// If package does not exists
			if count == 0 {
				database.Connection().Create(&pkg)
			} else {
				database.Connection().Model(&entities.Package{}).
					Where("name = ? and version = ?", pkg.Name, pkg.Version).First(&pkg)
			}
			assetPackages = append(assetPackages, pkg)
		}

		// Replace packages
		asset.Packages = assetPackages
		database.Connection().Model(&asset).Association("Packages").Replace(&assetPackages)
	}

	// Update discovery status
	discovery.UpdateStatus(entities.DiscoveryStatusDone, "Discovery completed successfully.")
}
