#!/bin/bash
# install_colo.sh - A script to download and install the 'colo' log colorizer tool.

# Exit immediately if a command exits with a non-zero status.
set -e

# Define variables for the binary and installation path
BINARY_URL="https://github.com/R-I-S-H-A-B-H-S-I-N-G-H/color-logs/raw/refs/heads/main/dist/colo"
INSTALL_DIR="/usr/local/bin"
BINARY_NAME="colo"
DEST_PATH="$INSTALL_DIR/$BINARY_NAME"

echo "Starting the installation of the '$BINARY_NAME' tool..."

# 1. Download the binary using curl
echo "Downloading '$BINARY_NAME' from GitHub..."
curl -L -o "$BINARY_NAME" "$BINARY_URL"

# 2. Make the binary executable
echo "Setting execute permissions..."
chmod +x "$BINARY_NAME"

# 3. Move the binary to a directory in the system's PATH
# This requires sudo because /usr/local/bin is a system directory.
echo "Installing '$BINARY_NAME' to $DEST_PATH (this requires admin rights)..."
sudo mv "$BINARY_NAME" "$DEST_PATH"

echo ""
echo "✅ Success! The '$BINARY_NAME' command is now installed."
echo "You can use it by piping any command's output to it."
echo "For example: tail -f your_log_file.log | $BINARY_NAME"
