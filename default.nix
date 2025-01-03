{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
, buildGoApplication ? pkgs.buildGoApplication
}:

buildGoApplication {
  pname = "ygot";
  version = "0.1";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;
  excludedPackages = [
    "./demo/getting_started"
    "./demo/uncompressed"
    "./integration_tests/uncompressed/uschema"
  ];
  doCheck = false;
}
