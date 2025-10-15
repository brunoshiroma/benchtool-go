#ANDROID_SDK=/SDK_HOME
#NDK_BIN=/workspaces/benchtool-go/android-ndk-r27d/toolchains/llvm/prebuilt/linux-x86_64/bin

build:
	go build cmd/benchtool-go/main.go

clean:
	go clean -modcache -cache -testcache -i -r
	rm *.h main lib

android-arm64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm64 \
	CC=$(NDK_BIN)/aarch64-linux-android31-clang \
	go build -gccgoflags "-Wl,-z,max-page-size=16384" -buildmode=c-shared -o benchtool-go-lib-android-arm64.so pkg/lib.go

android-armv7:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=arm \
	GOARM=7 \
	CC=$(NDK_BIN)/armv7a-linux-androideabi31-clang \
	go build -buildmode=c-shared -o benchtool-go-lib-android-armv7.so pkg/lib.go

android-386:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=386 \
	CC=$(NDK_BIN)/i686-linux-android31-clang \
	go build -buildmode=c-shared -o benchtool-go-lib-android-386.so pkg/lib.go

android-x64:
	CGO_ENABLED=1 \
	GOOS=android \
	GOARCH=amd64 \
	CC=$(NDK_BIN)/x86_64-linux-android31-clang \
	go build -buildmode=c-shared -o benchtool-go-lib-android-x64.so pkg/lib.go

android: android-x64 android-386 android-armv7 android-arm64