const hre = require("hardhat");
const { ethers, upgrades } = require("hardhat");
const { saveContractAddress, getSavedContractAddresses } = require("../utils");
const salesConfig = require("../sales_config_refresher");

async function main() {
  const network = hre.network.name;
  const config = require("../configs/saleConfig.json")[network];

  // ─── 1. C2N Token ───────────────────────────────────────────────
  console.log("\n>>> [1/7] Deploying C2N Token...");
  const C2N = await ethers.getContractFactory("C2NToken");
  const c2nToken = await C2N.deploy("C2N", "C2N", "1000000000000000000000000000", 18);
  await c2nToken.waitForDeployment();
  const c2nAddress = await c2nToken.getAddress();
  console.log("C2N deployed to:", c2nAddress);
  saveContractAddress(network, "C2N-TOKEN", c2nAddress);

  // ─── 2. Airdrop ──────────────────────────────────────────────────
  console.log("\n>>> [2/7] Deploying Airdrop...");
  const Airdrop = await ethers.getContractFactory("Airdrop");
  const airdrop = await Airdrop.deploy(c2nAddress);
  await airdrop.waitForDeployment();
  const airdropAddress = await airdrop.getAddress();
  console.log("Airdrop deployed to:", airdropAddress);
  saveContractAddress(network, "Airdrop-C2N", airdropAddress);
  let tx = await c2nToken.transfer(airdropAddress, ethers.parseEther("10000"));
  await tx.wait();
  tx = await airdrop.withdrawTokens();
  await tx.wait();
  console.log("Airdrop funded and tested.");

  // ─── 3. Farm ─────────────────────────────────────────────────────
  console.log("\n>>> [3/7] Deploying Farm...");
  const now = Math.round(Date.now() / 1000);
  const Farm = await ethers.getContractFactory("FarmingC2N");
  const farm = await Farm.deploy(c2nAddress, ethers.parseEther("1"), now + 100);
  await farm.waitForDeployment();
  const farmAddress = await farm.getAddress();
  console.log("Farm deployed to:", farmAddress);
  saveContractAddress(network, "FarmingC2N", farmAddress);
  tx = await c2nToken.approve(farmAddress, ethers.parseEther("50000"));
  await tx.wait();
  tx = await farm.fund(ethers.parseEther("50000"));
  await tx.wait();
  await farm.add(100, c2nAddress, true);
  console.log("Farm funded and LP token added.");

  // ─── 4. IDO 基础设施 ──────────────────────────────────────────────
  console.log("\n>>> [4/7] Deploying IDO infrastructure...");
  const Admin = await ethers.getContractFactory("Admin");
  const admin = await Admin.deploy(config.admins);
  await admin.waitForDeployment();
  const adminAddress = await admin.getAddress();
  console.log("Admin deployed to:", adminAddress);
  saveContractAddress(network, "Admin", adminAddress);

  const SalesFactory = await ethers.getContractFactory("SalesFactory");
  const salesFactory = await SalesFactory.deploy(adminAddress, "0x0000000000000000000000000000000000000000");
  await salesFactory.waitForDeployment();
  const salesFactoryAddress = await salesFactory.getAddress();
  console.log("SalesFactory deployed to:", salesFactoryAddress);
  saveContractAddress(network, "SalesFactory", salesFactoryAddress);

  const currentTs = (await ethers.provider.getBlock("latest")).timestamp;
  const AllocationStaking = await ethers.getContractFactory("AllocationStaking");
  const allocationStaking = await upgrades.deployProxy(
    AllocationStaking,
    [c2nAddress, ethers.parseEther(config.allocationStakingRPS), currentTs + config.delayBeforeStart, salesFactoryAddress],
    { unsafeAllow: ["delegatecall"] }
  );
  await allocationStaking.waitForDeployment();
  const stakingAddress = await allocationStaking.getAddress();
  console.log("AllocationStaking deployed to:", stakingAddress);
  saveContractAddress(network, "AllocationStakingProxy", stakingAddress);

  const proxyAdmin = await upgrades.erc1967.getAdminAddress(stakingAddress);
  saveContractAddress(network, "ProxyAdmin", proxyAdmin);

  await salesFactory.setAllocationStaking(stakingAddress);
  tx = await c2nToken.approve(stakingAddress, ethers.parseEther(config.initialRewardsAllocationStaking));
  await tx.wait();
  tx = await allocationStaking.add(100, c2nAddress, true);
  await tx.wait();
  await allocationStaking.fund(ethers.parseEther(config.initialRewardsAllocationStaking));
  console.log("AllocationStaking funded.");

  // ─── 5. MOCK Token ───────────────────────────────────────────────
  console.log("\n>>> [5/7] Deploying MOCK Token...");
  const MCK = await ethers.getContractFactory("C2NToken");
  const mckToken = await MCK.deploy("MOCK-TOKEN", "MCK", "1000000000000000000000000000", 18);
  await mckToken.waitForDeployment();
  const mckAddress = await mckToken.getAddress();
  console.log("MOCK-TOKEN deployed to:", mckAddress);
  saveContractAddress(network, "MOCK-TOKEN", mckAddress);

  // ─── 6. Sale ─────────────────────────────────────────────────────
  console.log("\n>>> [6/7] Creating Sale...");
  salesConfig.refreshSalesConfig(network);
  // 清除 require 缓存，确保读到刚写入的新时间戳
  delete require.cache[require.resolve("../configs/saleConfig.json")];
  const freshConfig = require("../configs/saleConfig.json")[network];

  tx = await salesFactory.deploySale();
  await tx.wait();
  const saleAddress = await salesFactory.getLastDeployedSale();
  console.log("Sale deployed to:", saleAddress);

  const sale = await ethers.getContractAt("C2NSale", saleAddress);
  const registrationStart = freshConfig.registrationStartAt;
  const registrationEnd = registrationStart + freshConfig.registrationLength;
  const saleStartTime = registrationEnd + freshConfig.delayBetweenRegistrationAndSale;
  const saleEndTime = saleStartTime + freshConfig.saleRoundLength;

  tx = await sale.setSaleParams(
    mckAddress,
    freshConfig.saleOwner,
    ethers.parseEther(freshConfig.tokenPriceInEth),
    ethers.parseEther(freshConfig.totalTokens),
    saleEndTime,
    freshConfig.TGE,
    freshConfig.portionVestingPrecision,
    ethers.parseEther(freshConfig.maxParticipation)
  );
  await tx.wait();
  tx = await sale.setRegistrationTime(registrationStart, registrationEnd);
  await tx.wait();
  tx = await sale.setSaleStart(saleStartTime);
  await tx.wait();
  tx = await sale.setVestingParams(freshConfig.unlockingTimes, freshConfig.portionPercents, freshConfig.maxVestingTimeShift);
  await tx.wait();
  console.log("Sale params set.");

  // ─── 7. TGE ──────────────────────────────────────────────────────
  console.log("\n>>> [7/7] TGE - depositing tokens...");
  tx = await mckToken.approve(saleAddress, ethers.parseEther(freshConfig.totalTokens));
  await tx.wait();
  tx = await sale.depositTokens();
  await tx.wait();
  console.log("IDO sale deposited.");

  // ─── 完成 ─────────────────────────────────────────────────────────
  console.log("\n✓ All contracts deployed successfully!");
  console.log("  C2N Token:         ", c2nAddress);
  console.log("  Airdrop:           ", airdropAddress);
  console.log("  Farm:              ", farmAddress);
  console.log("  Admin:             ", adminAddress);
  console.log("  SalesFactory:      ", salesFactoryAddress);
  console.log("  AllocationStaking: ", stakingAddress);
  console.log("  MOCK-TOKEN:        ", mckAddress);
  console.log("  Sale:              ", saleAddress);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
