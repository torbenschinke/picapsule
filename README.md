# picapsule

picapsule is a heritage time capsule which consists of a hardware part and a software part.
The hardware is based on a Raspberry Pi 4 or better and a Nago Web Application which is displayed on a Raspberry Touchdisplay 2 within a webbrowser in kiosk mode as UI.
Because it does not need any batteries or internet it has the potential to last decades and must only plugged with power and call it a day.

## install

install.sh script

```bash
#!/bin/sh

set -e

GOPROXY=direct go install github.com/torbenschinke/picapsule/cmd/main@latest

# Pfad zur installierten Binary
APP_BIN="$HOME/go/bin/main"

# Autostart-Verzeichnis erstellen
AUTOSTART_DIR="$HOME/.config/autostart"
mkdir -p "$AUTOSTART_DIR"

# .desktop-Datei für Autostart erstellen
cat > "$AUTOSTART_DIR/picapsule.desktop" <<EOF
[Desktop Entry]
Type=Application
Name=Picapsule
Exec=$APP_BIN
X-GNOME-Autostart-enabled=true
EOF

echo "Installation abgeschlossen. Die Anwendung startet nun automatisch beim Desktop-Login."

```

## fix firefox:
about:config im Profil (once manual

```
browser.sessionstore.resume_from_crash = false

browser.sessionstore.max_resumed_crashes = 0
```

