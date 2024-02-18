# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoCheckUpdates < Formula
  desc "go-check-updates upgrades your go.mod dependencies to the latest versions, ignoring specified versions."
  homepage "https://github.com/fogfish/go-check-updates"
  version "0.4.1"
  license "MIT"

  depends_on "go" => :optional

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.1/go-check-updates_0.4.1_darwin_arm64"
      sha256 "7cebdc8b6756cc59bac26369499916eeeec76d0b492ddd0f2bfecce328f7e003"

      def install
        bin.install "go-check-updates_0.4.1_darwin_arm64" => "go-check-updates"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.1/go-check-updates_0.4.1_darwin_amd64"
      sha256 "ed932b94bf0d0d09162c58f03b39fc9456e9c9be3b148c0c1d9c217b146b28cd"

      def install
        bin.install "go-check-updates_0.4.1_darwin_amd64" => "go-check-updates"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.1/go-check-updates_0.4.1_linux_amd64"
      sha256 "7b0d68fdc8a65f57cc869a74a425d1f82510165d1307c78338243df605443e8c"

      def install
        bin.install "go-check-updates_0.4.1_linux_amd64" => "go-check-updates"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.1/go-check-updates_0.4.1_linux_arm64"
      sha256 "5a8496d94333c55bbcf91cd80329b268e10bf7021d3a59a1cf37242971abfb28"

      def install
        bin.install "go-check-updates_0.4.1_linux_arm64" => "go-check-updates"
      end
    end
  end

  test do
    system "#{bin}/go-check-updates -v"
  end
end
