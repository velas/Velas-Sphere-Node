//TODO general
// Make the upgradable contract
// Check how ofter the node inside the majority. in case when it is offline for long time exclude it. it losses membershipFee

pragma solidity >=0.4.22 <0.6.0;

contract VelasSphere {
    uint membershipFee = 100000000000;
    // maybe temporary
    uint gracePeriod = 50;
    uint nodeCount;
    uint minNodesVoices;
    uint customerCount;
    uint poolCount;
    //TODO is it enough?
    uint votesToBanPermanently = 5;

    struct Pricing {
        uint keepPerByte;
        uint writePerByte;
        uint GPUTPerCycle;
        uint CPUTtPerCycle;
        bool isChanged; //was it changed by the customer;
    }

    Pricing defaultPricing; //default market pricing

    constructor() public {
        defaultPricing.keepPerByte = 1;
        defaultPricing.writePerByte = 1;
        defaultPricing.GPUTPerCycle = 1;
        defaultPricing.CPUTtPerCycle = 1;
        defaultPricing.isChanged = false;
        // 2/3 of 94 nodes. Need to be count?
        minNodesVoices = 62;
    }

    struct Location {
        uint pool; //Pool of 94 nodes joined together.
        uint place; //One of 94 position in that generation. Once all of 94 positon is busy in current generation, so need to move to next generation
    }

    struct Pool {
        uint poolID; // number of pool
        uint nodeCount;
    }

    mapping (uint => Pool) pools;

    struct Node {
        address payable staking_addr;
        address mining_addr;
        uint balance;
        bool active;
        Pricing pricing;
        Location location;
    }

    struct NodeToBan {
         address addr;
         uint votes;
    }

    mapping(address => Node) nodes;
    mapping(address => NodeToBan) banned;

    struct Customer {
        Pricing pricing;
        // customer choose which pool can create invoices
        bool specificPool;
        Location location;
        uint balance;
        bool registered;
        mapping(address => NodeToBan) banned;
    }

    mapping(address => Customer) customers;

    function() external payable {
        deposit();
    }

    function banNode(address _node) external {
         Customer storage current = customers[msg.sender];
         //it's enough in this customer's map
         current.banned[_node].votes = 1;
         banned[_node].votes += 1;

         if (banned[_node].votes == votesToBanPermanently) {
             //TODO ban permanently
         }
    }

    //Customer may want to increase the price to be first in list
    function proposePricing(uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) external {
        Customer storage current = customers[msg.sender];
        current.pricing.keepPerByte = _keepPerByte;
        current.pricing.writePerByte = _writePerByte;
        current.pricing.GPUTPerCycle = _GPUTPerCycle;
        current.pricing.CPUTtPerCycle = _CPUTtPerCycle;
        current.pricing.isChanged = true;
    }

    function deposit() internal {
        Customer storage current = customers[msg.sender];
        require(msg.value > 0);
        current.balance += msg.value;
        if (current.registered == false) {
        current.pricing = defaultPricing;
            customerCount += 1;
        }
        current.registered = true;
    }

    //pull - user can define a specific pool. if he defines 0 then all pools
    //_places - user can define a specific places in a pool. if 0 all places
    function depositWithNodes(uint _pull, uint _places) external payable {
        deposit();
        if (_pull == 0 && _places == 0) {
            return;
        }
        changePool(_pull, _places);
    }

    function changePool(uint _pull, uint _places) public {
        Customer storage current = customers[msg.sender];
        current.location.place = _places;
        current.location.pool = _pull;
        current.specificPool = true;
    }

    struct Invoice {
        uint height_start;
        uint height_end;
        uint deadline;
        bool isOpened;
        address user;
        Pricing pricing;
        Resources used;
        uint price;
        uint voices;
    }

    struct Resources {
        uint keepPerByte;
        uint writePerByte;
        uint GPUTPerCycle;
        uint CPUTtPerCycle;
    }

    mapping(address => Invoice) invoices;

    function openInvoice(address addr, uint deadline) external {
        //TODO check customer
        require(addr == msg.sender);
        Invoice storage invoice = invoices[addr];
        //TODO must be new?
        require(invoice.voices == 0);
        invoice.deadline = deadline;
        invoice.isOpened = true;
    }

    //node sends the invoice to decrease the balance of the customer
    function createInvoice(address user, uint price) external  {
            Node storage node = nodes[msg.sender];
            require(node.active);
            //TODO check if node was permanently banned
            require(banned[node.staking_addr].votes < votesToBanPermanently);
            Customer storage customer = customers[user];
            //check if user banned node
            require(customer.banned[node.staking_addr].votes == 0);

            Invoice storage invoice = invoices[user];
            require(invoice.isOpened);

            invoice.voices += 1;
            invoice.price += price;
    }

    function calculatePrice(Pricing storage pricing, uint keepPerByte, uint writePerByte, uint GPUTPerCycle, uint CPUTtPerCycle) internal returns (uint) {
        uint price;
        price += pricing.keepPerByte * keepPerByte;
        price += pricing.writePerByte * writePerByte;
        price += pricing.CPUTtPerCycle * CPUTtPerCycle;
        price += pricing.GPUTPerCycle * GPUTPerCycle;
        return price;
    }

    function closeInvoice(address user, uint price) internal {
        //Nodes cannot work more than user requested
        require(price <= customers[user].balance);
        customers[user].balance -= price;
        delete invoices[user];
    }

    function getNextBitPosition() internal returns (uint) {
        Pool storage pool = pools[poolCount];
        if (pool.nodeCount >= 94) {
               poolCount += 1;
        }

        uint position;
            position = 1 << pool.nodeCount;
            pool.nodeCount += 1;
           return position;

    }

    function isRegistered(address mining_addr) internal returns (bool) {
        return nodes[mining_addr].active;
    }

    function registerNode(address payable staking_addr, address mining_addr) external payable {
        Node storage node = nodes[mining_addr];
        //TODO need to check if it exists
        require(node.active == false);
        node.active = true;
        require(msg.value == membershipFee);

        node.staking_addr = staking_addr;
        node.location.place = getNextBitPosition();
        node.location.pool = poolCount;

        nodeCount += 1;
    }

    function changeNodePricing(uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) external {
        Node storage node = nodes[msg.sender];
        require(node.active == true);
        node.pricing.keepPerByte = _keepPerByte;
        node.pricing.writePerByte = _writePerByte;
        node.pricing.GPUTPerCycle = _GPUTPerCycle;
        node.pricing.CPUTtPerCycle = _CPUTtPerCycle;
    }

    function withdraw(address payable addr) external {
        //TODO check staking signature
        Node storage node = nodes[addr];
        require(node.balance > 0);
        node.staking_addr.transfer(nodes[addr].balance);
        node.balance = 0;
    }

    function changeMiningAddr(address old_addr, address new_addr) external {
        //TODO check staking signature
        Node storage node = nodes[old_addr];
        node.mining_addr = new_addr;
    }
}
