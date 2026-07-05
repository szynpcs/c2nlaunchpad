$commands = @(
    "npx hardhat compile",
    "npx hardhat run --network local scripts/deployment/deploy_c2n_token.js",
    "npx hardhat run --network local scripts/deployment/deploy_airdrop_c2n.js",
    "npx hardhat run --network local scripts/deployment/deploy_farm.js",
    "npx hardhat run --network local scripts/deployment/deploy_ido.js",
    "npx hardhat run --network local scripts/deployment/deploy_sales_token.js",
    "npx hardhat run --network local scripts/deployment/deploy_sales.js",
    "npx hardhat run --network local scripts/deployment/deploy_tge.js"
)

foreach ($cmd in $commands) {
    Write-Host "`n>>> $cmd" -ForegroundColor Cyan
    Invoke-Expression $cmd
    if ($LASTEXITCODE -ne 0) {
        Write-Host "Failed at: $cmd" -ForegroundColor Red
        exit 1
    }
}

Write-Host "`nAll contracts deployed successfully!" -ForegroundColor Green
