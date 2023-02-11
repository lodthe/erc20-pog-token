import { ethers } from "hardhat";

async function main() {
  const initialSupply = '10000000000000000000000'; // 10e4 * 1e18
  const POGToken = await ethers.getContractFactory("POGToken");
  
  console.log('Deploying POGToken...');

  const token = await POGToken.deploy(initialSupply);

  await token.deployed();

  console.log(`POGToken has been deployed to ${token.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
