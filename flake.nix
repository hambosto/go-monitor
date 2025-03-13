{
  description = "Modern terminal-based system monitoring with real-time metrics and elegant TUI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:
    let
      nixosModule =
        {
          config,
          lib,
          pkgs,
          ...
        }:
        let
          cfg = config.programs.go-monitor;
        in
        {
          options.programs.go-monitor = {
            enable = lib.mkEnableOption "Enable the Go Monitor";
          };

          config = lib.mkIf cfg.enable {
            home.packages = [ self.packages.${pkgs.system}.default ];
          };
        };

      perSystem =
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};

          package = pkgs.buildGoModule {
            pname = "go-monitor";
            version = "1.0";

            src = ./.;

            vendorHash = "sha256-ehheKIkK8aBftGETdK73HntHfFBGjjdDMmyOzCLsyEo=";

            env.CGO_ENABLED = 0;

            ldflags = [
              "-extldflags '-static'"
              "-s -w"
            ];
          };
        in
        {
          packages.default = package;
        };
    in
    flake-utils.lib.eachDefaultSystem perSystem
    // {
      nixosModules.default = nixosModule;
    };
}
