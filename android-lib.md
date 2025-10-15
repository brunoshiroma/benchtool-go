# Compiling to android +35 api

Android api +35 requires 16kb page files

```bash
# download r27 ndk https://developer.android.com/ndk/downloads
# the last r29, generate a arm64 broken, crashing the app on lib load

wget https://dl.google.com/android/repository/android-ndk-r27-linux.zip
unzip android-ndk-r27-linux.zip
export NDK_BIN=/workspaces/benchtool-go/android-ndk-r27d/toolchains/llvm/prebuilt/linux-x86_64/bin

make build-android
```