# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class GoCheckUpdates < Formula
  desc "go-check-updates upgrades your go.mod dependencies to the latest versions, ignoring specified versions."
  homepage "https://github.com/fogfish/go-check-updates"
  version "0.4.4"
  license "MIT"

  depends_on "go" => :optional

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.4/go-check-updates_0.4.4_darwin_amd64"
      sha256 "bcec563fce2728e55144e09b0d3c604ab04d80872f0368af9e291b79f3f5cf5a"

      def install
        bin.install "go-check-updates_0.4.4_darwin_amd64" => "go-check-updates"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.4/go-check-updates_0.4.4_darwin_arm64"
      sha256 "30aeb80fae5d3e1f19c2f7745b2947827fbaba6c0f4d678c004f69373b01c952"

      def install
        bin.install "go-check-updates_0.4.4_darwin_arm64" => "go-check-updates"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.4/go-check-updates_0.4.4_linux_amd64"
      sha256 "31ae9306dc93e6ffca73167744ee3a3a0caf2b3a32f611042e6edb0e25c085aa"

      def install
        bin.install "go-check-updates_0.4.4_linux_amd64" => "go-check-updates"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/fogfish/go-check-updates/releases/download/v0.4.4/go-check-updates_0.4.4_linux_arm64"
      sha256 "8ad5ae0218767c28d079dc4f2263bc46347a0522aaf85ff3735a1f17f6627039"

      def install
        bin.install "go-check-updates_0.4.4_linux_arm64" => "go-check-updates"
      end
    end
  end

  test do
    system "#{bin}/go-check-updates -v"
  end
end
