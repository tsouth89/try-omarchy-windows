# v1 readiness

The Windows app and the Mac app now live under Omacom. A stable Windows release
still needs the checks below; the open work is tracked in one place, GitHub
issue #77. A passing build does not establish hardware or
upgrade reliability.

## Runtime and release validation

- [ ] Validate the source-built runtime on physical Windows hardware, including full Hyper-V. Record archive hashes and results using
  [RUNTIME-VALIDATION.md](RUNTIME-VALIDATION.md).
- [ ] Pin the tested runtime and matching source archive in
  `guest-build/runtime.lock.json`, then verify the runtime packaged in the
  release candidate.
- [ ] Verify signed draft preparation and publication using
  [RELEASING.md](RELEASING.md). A successful signing check alone does not
  establish that the full release workflow works.
- [ ] Test a copied pre-transfer installation through the signed update path
  into the current candidate. Verify redirects, preserved files, and rollback
 , separately from preview-to-stable migration.
- [ ] Prepare release notes and verify public download links, signatures, and
  versions after publication.

## Merged for release testing

- [#34](https://github.com/omacom/try-omarchy-windows/pull/34): stable updates,
  including a bridge for installations that skip preview releases.
- [#35](https://github.com/omacom/try-omarchy-windows/pull/35): disk capacity,
  in-place growth, and Windows free-space information.
- [#29](https://github.com/omacom/try-omarchy-windows/pull/29): official artwork,
  current resources, and splash icon handling.
- First-run install-location selection is on master and needs release testing.

## Release gates

- [x] Reproduce and resolve the configuration report in #32, or document its
  confirmed cause and supported fix. Fixed in the v0.0.13 image (guest patch 0033).
- [ ] Test the signed candidate on the Windows hardware matrix, including
  the source runtime, full Hyper-V, and remote input.
- [ ] Record fresh install, existing guest upgrade, interrupted update, and
  forced rollback. Include an old preview that skips the bridge and
  reaches stable through both signed update feeds.
- [ ] Verify disk growth inside the guest, unchanged user files, and unchanged
  capacity after lowering the setting or rolling back the launcher.
- [ ] Restore a configuration export onto a fresh physical Omarchy install.
  Confirm package and theme restoration and exclusion of VM-specific state.
- [x] Finish stopped-VM backup/restore and a clear reset flow before presenting
  the guest as suitable for persistent work. Shipped in v0.0.12. Configuration export is not a VM
  backup.
- [ ] Exercise sleep/resume, mixed-DPI resize, audio-device changes, and a long
  session. Record idle CPU and whether launch alone activates the microphone.
- [ ] Update user documentation to describe the tested release, including how
  existing guests receive Omarchy OS updates separately from launcher updates.

Use [TESTING.md](TESTING.md) for reports. Each gate needs evidence for the exact
candidate version and runtime, not only an earlier preview. Keep release
publication separate from code review and merging.

## Scope

v1 should provide a dependable way to try Omarchy, keep a trial setup, and take
its configuration to a full installation. Prioritize reliability, storage,
recovery, and understandable controls.

Image clipboard and better file transfers can follow the core work. Camera
bridging, Windows ARM64, and additional portable launchers remain later work.
Booting an existing physical installation is outside the v1 scope.
