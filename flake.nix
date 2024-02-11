{
  description = "Exercise on data grouping by time interval";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-23.05";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        pname = "data-grouping-exercise";
        version = "0.0.1";
      in rec {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.wgo
            pkgs.semgrep
            pkgs.gopls
          ];
          shellHook = ''
            export GOPATH=$PWD/.go
            export PATH=$GOPATH/bin:$PATH
          '';
        };
      }
    );
}
