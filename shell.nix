{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [

    git

    # Go development
    go_1_20
    gopls
    godef
    gotools

  ];
}
