{
  description = "crit work environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {inherit system;};
    in {
      devShells = {
        default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go_1_22
            gotools
            golangci-lint
            gopls
            go-outline
            gopkgs
            gofumpt
            gnumake
          ];
        };
      };
    });
}
