// Sources have been taken from 
// https://yuichiroaoki.medium.com/testing-erc20-smart-contracts-in-typescript-hardhat-9ad20eb40502

import { expect } from 'chai';
import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { POGToken__factory, POGToken } from '../typechain-types';

describe('Token contract', function () {

	let POGToken: POGToken;
	let owner: SignerWithAddress;
	let addr1: SignerWithAddress;
	let addr2: SignerWithAddress
	let addrs: SignerWithAddress[];

	beforeEach(async function () {
		[owner, addr1, addr2, ...addrs] = await ethers.getSigners();

		const POGTokenFactory = (await ethers.getContractFactory(
			'POGToken', owner
		)) as POGToken__factory;
		const totalSupply = (10 ** 9).toString()
		POGToken = await POGTokenFactory.deploy(
			ethers.utils.parseEther(totalSupply),
		)
	});

	describe('Deployment', function () {
		it('Should assign the total supply of tokens to the owner', async function () {
			const ownerBalance = await POGToken.balanceOf(owner.address);
			expect(await POGToken.totalSupply()).to.equal(ownerBalance);
		});
	});

	describe('Transactions', function () {
		it('Should transfer tokens between accounts', async function () {
			// Transfer 50 tokens from owner to addr1
			await POGToken.transfer(addr1.address, 50);
			const addr1Balance = await POGToken.balanceOf(addr1.address);
			expect(addr1Balance).to.equal(50);

			// Transfer 50 tokens from addr1 to addr2
			// We use .connect(signer) to send a transaction from another account
			await POGToken.connect(addr1).transfer(addr2.address, 50);
			const addr2Balance = await POGToken.balanceOf(addr2.address);
			expect(addr2Balance).to.equal(50);
		});

		it('Should fail if sender doesn’t have enough tokens', async function () {
			const initialOwnerBalance = await POGToken.balanceOf(owner.address);

			// Try to send 1 token from addr1 (0 tokens) to owner (1000 tokens).
			// `require` will evaluate false and revert the transaction.
			await expect(
				POGToken.connect(addr1).transfer(owner.address, 1)
			).to.be.revertedWith('ERC20: transfer amount exceeds balance');

			// Owner balance shouldn't have changed.
			expect(await POGToken.balanceOf(owner.address)).to.equal(
				initialOwnerBalance
			);
		});

		it('Should update balances after transfers', async function () {
			const initialOwnerBalance = await POGToken.balanceOf(owner.address);

			// Transfer 100 tokens from owner to addr1.
			await POGToken.transfer(addr1.address, 100);

			// Transfer another 50 tokens from owner to addr2.
			await POGToken.transfer(addr2.address, 50);

			// Check balances.
			const finalOwnerBalance = await POGToken.balanceOf(owner.address);
			expect(finalOwnerBalance).to.equal(initialOwnerBalance.sub(150));

			const addr1Balance = await POGToken.balanceOf(addr1.address);
			expect(addr1Balance).to.equal(100);

			const addr2Balance = await POGToken.balanceOf(addr2.address);
			expect(addr2Balance).to.equal(50);
		});
	});
});