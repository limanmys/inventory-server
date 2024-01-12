Name: inventory-server
Version: %VERSION%
Release: 0
License: MIT
Prefix: /opt
Summary: inventory server.
Group: Applications/System
BuildArch: x86_64

%description
inventory server.

%pre

%prep

%build

%install
cp -rfa %{_app_dir} %{buildroot}

%post -p /bin/bash
if [ -f "/usr/lib/systemd/system/inventory-server.service" ]; then
    rm -rf /usr/lib/systemd/system/inventory-server.service  2>/dev/null || true
    systemctl disable inventory-server.service 2>/dev/null || true
    systemctl stop inventory-server.service 2>/dev/null || true
    systemctl daemon-reload 2>/dev/null || true
fi

echo """
[Unit]
Description=Inventory Server (%I)
[Service]
Type=simple
WorkingDirectory=/opt/inventory-server
ExecStart=/opt/inventory-server/inventory-server -type=%i
Restart=always
RestartSec=10
SyslogIdentifier=inventory
KillSignal=SIGINT
User=root
Group=root
[Install]
WantedBy=multi-user.target
    """ > /etc/systemd/system/inventory-server@.service

systemctl daemon-reload
systemctl enable inventory-server@admin.service
systemctl restart inventory-server@admin.service

%clean

%files
%defattr(0770, root, root)
/opt/inventory-server/*
/opt/inventory-server/.env.example

%define _unpackaged_files_terminate_build 0