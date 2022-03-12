// SPDX-License-Identifier: MIT

pragma solidity >=0.5.0;


contract helloworld {

	mapping (address => uint256 )  balance;
	
    function say() public pure returns (string memory) {
    	
        return 'hello etherworld';
    }
    
    
    function deposit(uint256 amount) public {
    	
		balance[msg.sender] += amount;
        
    }

    function checkBalance(address account) public view returns(uint256){
    	
		return balance[account];
        
    }

    function withdraw(uint256 amount) public {
    	
		if ( balance[msg.sender] >  amount ) {
			balance[msg.sender] -= amount ;
		}
		
        
    }

}