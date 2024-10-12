{
  description = "Crossplane provider for Cloudflare";

  # Nixpkgs / NixOS version to use.
  # inputs.nixpkgs.url = "nixpkgs/nixos-24.11";
  # TODO: nixos-24.11 is not currently available, and required to build some
  #       flake, so we will use this one too too avoid using two different
  #       nixpkgs versions and increasing the space usage.
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs =
    { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };
      in
      rec {
        devShells.default =
          pkgs.mkShell {
            packages = [
              pkgs.commitlint
              pkgs.coreutils
              pkgs.direnv
              pkgs.dogdns
              pkgs.git
              pkgs.gnumake
              pkgs.go
              pkgs.golangci-lint
              pkgs.helm
              pkgs.jq
              pkgs.kind
              pkgs.kubectl
              pkgs.lefthook
              pkgs.trunk-io
              pkgs.yq-go
              pkgs.upbound
            ];

            installPhase = "";
          };
      }
    );
}
