#!/bin/bash
# ===================================================
# Qoo10JP Order Go - Ubuntu Server Setup Script
# ===================================================

set -e

APP_NAME="qoo10jp-order-go"
APP_DIR="/home/james/qoo10jp-order-go"
SERVICE_FILE="/etc/systemd/system/${APP_NAME}.service"

echo "=========================================="
echo "Qoo10JP Order Go - Ubuntu Setup"
echo "=========================================="

# 1. Create app directory
echo "[1/5] Creating app directory..."
mkdir -p $APP_DIR
mkdir -p $APP_DIR/web

# 2. Copy service file
echo "[2/5] Installing systemd service..."
if [ -f "./qoo10jp-order-go.service" ]; then
    sudo cp ./qoo10jp-order-go.service $SERVICE_FILE
    sudo systemctl daemon-reload
    echo "  Service file installed: $SERVICE_FILE"
else
    echo "  Warning: Service file not found, skipping..."
fi

# 3. Set permissions
echo "[3/5] Setting permissions..."
chmod +x $APP_DIR/${APP_NAME} 2>/dev/null || true
chown -R james:james $APP_DIR

# 4. Enable and start service
echo "[4/5] Enabling service..."
sudo systemctl enable $APP_NAME 2>/dev/null || true

# 5. Display status
echo "[5/5] Setup complete!"
echo ""
echo "=========================================="
echo "Available commands:"
echo "=========================================="
echo "  sudo systemctl start $APP_NAME   - Start service"
echo "  sudo systemctl stop $APP_NAME    - Stop service"
echo "  sudo systemctl restart $APP_NAME - Restart service"
echo "  sudo systemctl status $APP_NAME  - Check status"
echo "  journalctl -u $APP_NAME -f       - View logs"
echo ""
echo "  Or use manual control:"
echo "  cd $APP_DIR && ./start.sh"
echo ""
echo "Ports:"
echo "  Shopee Order Go:   8080"
echo "  Qoo10JP Order Go:  8081"
echo "=========================================="
