{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [ pkgs.git pkgs.vscodium ];

  shellHook = ''
    codium --install-extension jnoortheen.nix-ide --force
    codium --install-extension naumovs.color-highlight --force 
    codium --install-extension golang.Go --force 
  '';
}
