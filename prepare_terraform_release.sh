#!/bin/bash

PROVIDER_NAME="terraform-provider-solacebroker"

# Build the provider for multiple platforms and architectures
function build_provider() {
    local os=$1
    local arch=$2
    local output_dir="bin/${os}_${arch}"

    echo "Building $PROVIDER_NAME $PROVIDER_VERSION for $os $arch..."
    local filename="${PROVIDER_NAME}_${PROVIDER_VERSION}_${os}_${arch}"
    GOOS=$os GOARCH=$arch go build -o "$output_dir/$filename"
}

mkdir -p bin
mkdir -p dist

PLATFORMS=("windows" "linux" "darwin" "freebsd")
ARCHITECTURES=("amd64" "arm64" "386" "arm")

# Loop through each platform and architecture and build the provider
for os in "${PLATFORMS[@]}"; do
    for arch in "${ARCHITECTURES[@]}"; do
        build_provider "$os" "$arch"
    done
done

SHASUMS_FILE="../dist/${PROVIDER_NAME}_${PROVIDER_VERSION}_SHA256SUMS"
SHASUMS_SIG_FILE="${SHASUMS_FILE}.sig"

# Compress binaries into zip files and generate shasums
cd bin
for os in "${PLATFORMS[@]}"; do
    for arch in "${ARCHITECTURES[@]}"; do
        file="${os}_${arch}/${PROVIDER_NAME}_${PROVIDER_VERSION}_${os}_${arch}"
        if [[ -f "$file" ]]; then
            echo "Creating zip for $file..."
            tar -czf "../dist/${PROVIDER_NAME}_${PROVIDER_VERSION}_${os}_${arch}.zip" "$file"
            cd ../dist
            sha256sum "${PROVIDER_NAME}_${PROVIDER_VERSION}_${os}_${arch}.zip" >> "$SHASUMS_FILE"
            cd ../bin
        fi
    done
done

#Sign file
gpg --batch --pinentry-mode loopback --passphrase ${GPG_PASSPHRASE} --armor --detach-sign --output ${SHASUMS_SIG_FILE} ${SHASUMS_FILE}

cd ..

# Copy and rename the terraform-registry-manifest.json file
MANIFEST_SOURCE="terraform-registry-manifest.json"
MANIFEST_DEST="dist/${PROVIDER_NAME}-${PROVIDER_VERSION}_manifest.json"
cp "$MANIFEST_SOURCE" "$MANIFEST_DEST"

echo "Release preparation completed."
