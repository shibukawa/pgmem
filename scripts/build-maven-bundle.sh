#!/usr/bin/env bash
# Builds the Maven Central bundle dist/maven/pgmem-<version>-bundle.zip:
# io.github.shibukawa.pgmem:pgmem and pgmem-junit5 (jar, sources, javadoc) and pgmem-native
# (one classifier jar per platform from dist/<goos>-<goarch>/), each file
# with its POM, checksums and signature. GPG_PRIVATE_KEY (an armored secret
# key) and GPG_PASSPHRASE sign it; the version is the one in
# packages/java/gradle.properties (scripts/set-version.sh).
#
#   scripts/build-binaries.sh && scripts/build-maven-bundle.sh
set -euo pipefail
cd "$(dirname "$0")/.."
root=$PWD
: "${GPG_PRIVATE_KEY:?set GPG_PRIVATE_KEY (armored secret key) to sign the bundle}"

version=$(sed -n 's/^version=//p' packages/java/gradle.properties)
staging=packages/java/build/central-staging
rm -rf "$staging" dist/maven
(cd packages/java && ./gradlew --no-daemon --console=plain -PnativeDist="$root/dist" \
  publishAllPublicationsToCentralStagingRepository)

status=0
while IFS= read -r f; do
  for ext in asc md5 sha1; do
    if [ ! -f "$f.$ext" ]; then echo "missing $f.$ext"; status=1; fi
  done
done < <(find "$staging" -type f \( -name '*.jar' -o -name '*.pom' -o -name '*.module' \))
for classifier in darwin-arm64 darwin-x86_64 linux-arm64 linux-x86_64 windows-arm64 windows-x86_64; do
  jar=$staging/io/github/shibukawa/pgmem/pgmem-native/$version/pgmem-native-$version-$classifier.jar
  if [ ! -f "$jar" ]; then echo "missing $jar"; status=1; fi
done
[ $status = 0 ]

mkdir -p dist/maven
(cd "$staging" && zip -qr "$root/dist/maven/pgmem-$version-bundle.zip" io -x '*maven-metadata.xml*')
unzip -l "dist/maven/pgmem-$version-bundle.zip" | grep -v '\.\(md5\|sha1\|sha256\|sha512\|asc\)$'
