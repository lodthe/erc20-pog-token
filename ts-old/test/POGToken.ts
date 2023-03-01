// Sources have been taken from 
// https://yuichiroaoki.medium.com/testing-erc20-smart-contracts-in-typescript-hardhat-9ad20eb40502

import { expect } from 'chai';
import { ethers } from 'hardhat';
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import { POGToken__factory, POGToken } from '../typechain-types';

describe('Token contract', function () {

	let POGToken: POGToken;
	let owner: SignerWithAddress;
	let addrAlice: SignerWithAddress;
	let addrBob: SignerWithAddress
	let addrs: SignerWithAddress[];

	beforeEach(async function () {
		[owner, addrAlice, addrBob, ...addrs] = await ethers.getSigners();

		const POGTokenFactory = (await ethers.getContractFactory(
			'POGToken', owner
		)) as POGToken__factory;
		const initialSupply = (10 ** 9).toString();

		POGToken = await POGTokenFactory.deploy(
			ethers.utils.parseEther(initialSupply),
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
			await POGToken.transfer(addrAlice.address, 50);
			const aliceBalance = await POGToken.balanceOf(addrAlice.address);
			expect(aliceBalance).to.equal(50);

			// Transfer 50 tokens from addr1 to addr2
			// We use .connect(signer) to send a transaction from another account
			await POGToken.connect(addrAlice).transfer(addrBob.address, 50);
			const bobBalance = await POGToken.balanceOf(addrBob.address);
			expect(bobBalance).to.equal(50);
		});

		it('Should fail if sender doesn’t have enough tokens', async function () {
			const initialOwnerBalance = await POGToken.balanceOf(owner.address);

			// Try to send 1 token from addr1 (0 tokens) to owner (1000 tokens).
			// `require` will evaluate false and revert the transaction.
			await expect(
				POGToken.connect(addrAlice).transfer(owner.address, 1)
			).to.be.revertedWith('ERC20: transfer amount exceeds balance');

			// Owner balance shouldn't have changed.
			expect(await POGToken.balanceOf(owner.address)).to.equal(
				initialOwnerBalance
			);
		});

		it('Should update balances after transfers', async function () {
			const initialOwnerBalance = await POGToken.balanceOf(owner.address);

			// Transfer 100 tokens from owner to addr1.
			await POGToken.transfer(addrAlice.address, 100);

			// Transfer another 50 tokens from owner to addr2.
			await POGToken.transfer(addrBob.address, 50);

			// Check balances.
			const finalOwnerBalance = await POGToken.balanceOf(owner.address);
			expect(finalOwnerBalance).to.equal(initialOwnerBalance.sub(150));

			const aliceBalance = await POGToken.balanceOf(addrAlice.address);
			expect(aliceBalance).to.equal(100);

			const bobBalance = await POGToken.balanceOf(addrBob.address);
			expect(bobBalance).to.equal(50);
		});
	});

  describe('Changing supply', function () {
        describe('Minting new tokens', function() {
      it('Minting new tokens', async function () {
        // Mint new tokens and validate balance and supply.
        const initialBalance = await POGToken.balanceOf(addrAlice.address);
        const initialTotalSupply = await POGToken.totalSupply();
        const mintedAmount = 50;
  
        await POGToken.mint(addrAlice.address, mintedAmount);
        const actualBalance = await POGToken.balanceOf(addrAlice.address);
        const actualTotalSupply = await POGToken.totalSupply();
  
        expect(actualBalance).to.equal(initialBalance.add(mintedAmount));
        expect(actualTotalSupply).to.equal(initialTotalSupply.add(mintedAmount));
      });
    });

    describe('Burning tokens', function() {
      it('Should success', async function () {
        const initialBalance = await POGToken.balanceOf(owner.address);
        const initialTotalSupply = await POGToken.totalSupply();
  
        // Check if balance if the balance is sufficient.
        const amountToBurn = 137;
        expect(initialBalance).to.greaterThan(amountToBurn);
  
        await POGToken.burn(owner.address, amountToBurn);
        const actualBalance = await POGToken.balanceOf(owner.address);
        const actualTotalSupply = await POGToken.totalSupply();
  
        expect(actualBalance).to.equal(initialBalance.sub(amountToBurn));
        expect(actualTotalSupply).to.equal(initialTotalSupply.sub(amountToBurn));
      });
  
      it('Should fail due to insufficient balance', async function () {
        const initialBalance = await POGToken.balanceOf(owner.address);
  
        // Try to burn x2 balance.
        const amountToBurn = initialBalance.mul(2);
        expect(initialBalance).to.lessThan(amountToBurn);
  
        await expect(
          POGToken.burn(owner.address, amountToBurn)
        ).to.be.revertedWith('ERC20: burn amount exceeds balance');
      });
    });
	});

    describe('Viewers', function() {
        describe('Adding', function() {
            it('Success', async function () {
                await expect(
                    POGToken.addViewer(addrAlice.address, 'alice', true)
                ).to.emit(POGToken, "LogViewerAdded").withArgs(addrAlice.address, 'alice');
            });

            it('Failure, already exists', async function () {
                await expect(
                    POGToken.addViewer(addrAlice.address, 'alice', true)
                ).to.emit(POGToken, "LogViewerAdded").withArgs(addrAlice.address, 'alice');

                await expect(
                    POGToken.addViewer(addrAlice.address, 'alice', true)
                ).to.be.revertedWith('A viewer with the given address has already been added');
            });
        });

        describe('Removing', function() {
            it('Success', async function () {
                await expect(
                    POGToken.addViewer(addrAlice.address, 'alice', true)
                ).to.emit(POGToken, "LogViewerAdded").withArgs(addrAlice.address, 'alice');

                await expect(
                    POGToken.removeViewer(addrAlice.address)
                ).to.emit(POGToken, "LogViewerRemoved").withArgs(addrAlice.address, 'alice');
            });

            it('Failure, not found', async function () {
                await expect(
                    POGToken.removeViewer(addrAlice.address)
                ).to.be.revertedWith('Viewer not found');
            });
        });
    });
});