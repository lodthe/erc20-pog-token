import { ethers } from "hardhat";

async function main() {
  const contractAddress = '0x8aADCf1ba17cFbD3D8b7A6230D1B2CFBE658d24D';
  const receiverAddress = '0xB87e1B6A462660E26Ec8913513aB86b084A4F46c';
  const mintAmount = ethers.utils.parseEther('1000');

  const POGToken = await (await ethers.getContractFactory("POGToken")).attach(contractAddress);
  const mint = await POGToken.mint(receiverAddress, mintAmount);
  
  console.log(`Successfully minted ${mintAmount} tokens, tx: ${mint.hash}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
