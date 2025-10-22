# BluePasser

PassPush is a smart Wi-Fi password management tool designed for offices and small teams. With PassPush, updating your network password is as simple as changing it on a single device — all other authorized devices automatically receive the new credentials via secure Bluetooth, saving time and eliminating the hassle of manual updates.

## Communication Flow

Admin device: password changed manually or via GUI.

Go core: encrypts the password and broadcasts via BLE.

Other devices: Go daemon listens for BLE updates.

C++ helper: writes the new password to the OS Wi-Fi manager.

Result: all devices are updated automatically, securely.

## Advantages of Go + C++

Cross-platform support: Go handles most logic, so your code runs on Windows, Linux, macOS.

System integration: C++ handles tasks that need deep OS-level access.

Performance: Both languages are fast; Go handles networking concurrency, C++ handles low-level updates.

Single binary deployment: Go builds are easy; C++ helpers can be bundled as dynamic libraries.
