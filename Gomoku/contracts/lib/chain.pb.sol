pragma solidity ^0.4.0;
import "./runtime.sol";


library pbRpcStateProof{
  //struct definition
  struct Data {   
    uint256 nonce;
    bytes state;
    bytes32 pendingConditionRoot;
    uint256 stateChannelId;
    uint256 maxCondTimeout;             
  }                           
  // Decoder section                       
  function decode(bytes bs) internal constant returns (Data) {
    var (x,) = _decode(32, bs, bs.length);                       
    return x;                                                    
  }
  function decode(Data storage self, bytes bs) internal constant {
    var (x,) = _decode(32, bs, bs.length);                    
    store(x, self);                                           
  }                             
  // innter decoder                       
  function _decode(uint p, bytes bs, uint sz)                   
      internal constant returns (Data, uint) {             
    Data memory r;                                          
    uint[6] memory counters;                                  
    uint fieldId;                                               
    _pb.WireType wireType;                                      
    uint bytesRead;                                             
    uint offset = p;                                            
    while(p < offset+sz) {                                     
      (fieldId, wireType, bytesRead) = _pb._decode_key(p, bs);  
      p += bytesRead;                                           
      if (false) {}
      else if(fieldId == 1)       
          p += _read_nonce(p, bs, r, counters);
      else if(fieldId == 2)       
          p += _read_state(p, bs, r, counters);
      else if(fieldId == 3)       
          p += _read_pendingConditionRoot(p, bs, r, counters);
      else if(fieldId == 4)       
          p += _read_stateChannelId(p, bs, r, counters);
      else if(fieldId == 5)       
          p += _read_maxCondTimeout(p, bs, r, counters);
      else revert();                                              
    }                                                          
    p = offset;                                                 
                                                    
    while(p < offset+sz) {                                     
      (fieldId, wireType, bytesRead) = _pb._decode_key(p, bs);  
      p += bytesRead;                                           
      if (false) {}
      else if(fieldId == 1)       
          p += _read_nonce(p, bs, nil(), counters);
      else if(fieldId == 2)       
          p += _read_state(p, bs, nil(), counters);
      else if(fieldId == 3)       
          p += _read_pendingConditionRoot(p, bs, nil(), counters);
      else if(fieldId == 4)       
          p += _read_stateChannelId(p, bs, nil(), counters);
      else if(fieldId == 5)       
          p += _read_maxCondTimeout(p, bs, nil(), counters);
      else revert();                                             
    }                                                          
    return (r, sz);                                             
  }                                                            
                            
  // field readers                       
  function _read_nonce(uint p, bytes bs, Data r, uint[6] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_uint256(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[1] += 1;                                            
    } else {                                                         
      r.nonce = x;                                         
      if(counters[1] > 0) counters[1] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_state(uint p, bytes bs, Data r, uint[6] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_bytes(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[2] += 1;                                            
    } else {                                                         
      r.state = x;                                         
      if(counters[2] > 0) counters[2] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_pendingConditionRoot(uint p, bytes bs, Data r, uint[6] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_bytes32(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[3] += 1;                                            
    } else {                                                         
      r.pendingConditionRoot = x;                                         
      if(counters[3] > 0) counters[3] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_stateChannelId(uint p, bytes bs, Data r, uint[6] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_uint256(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[4] += 1;                                            
    } else {                                                         
      r.stateChannelId = x;                                         
      if(counters[4] > 0) counters[4] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_maxCondTimeout(uint p, bytes bs, Data r, uint[6] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_uint256(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[5] += 1;                                            
    } else {                                                         
      r.maxCondTimeout = x;                                         
      if(counters[5] > 0) counters[5] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
                            
  // struct decoder                       
                                      
            
  //store function                                                     
  function store(Data memory input, Data storage output) internal{
    output.nonce = input.nonce;                           
    output.state = input.state;                           
    output.pendingConditionRoot = input.pendingConditionRoot;                           
    output.stateChannelId = input.stateChannelId;                           
    output.maxCondTimeout = input.maxCondTimeout;                           
  }                                                                   
             
  //utility functions                                           
  function nil() internal constant returns (Data r) {        
    assembly { r := 0 }                                       
  }                                                            
  function isNil(Data x) internal constant returns (bool r) {
    assembly { r := iszero(x) }                               
  }                                                            
} 

library pbRpcMultiSignature{
  //struct definition
  struct Data {   
    uint8[] v;
    bytes32[] r;
    bytes32[] s;             
  }                           
  // Decoder section                       
  function decode(bytes bs) internal constant returns (Data) {
    var (x,) = _decode(32, bs, bs.length);                       
    return x;                                                    
  }
  function decode(Data storage self, bytes bs) internal constant {
    var (x,) = _decode(32, bs, bs.length);                    
    store(x, self);                                           
  }                             
  // innter decoder                       
  function _decode(uint p, bytes bs, uint sz)                   
      internal constant returns (Data, uint) {             
    Data memory r;                                          
    uint[4] memory counters;                                  
    uint fieldId;                                               
    _pb.WireType wireType;                                      
    uint bytesRead;                                             
    uint offset = p;                                            
    while(p < offset+sz) {                                     
      (fieldId, wireType, bytesRead) = _pb._decode_key(p, bs);  
      p += bytesRead;                                           
      if (false) {}
      else if(fieldId == 1)       
          p += _read_v(p, bs, nil(), counters);
      else if(fieldId == 2)       
          p += _read_r(p, bs, nil(), counters);
      else if(fieldId == 3)       
          p += _read_s(p, bs, nil(), counters);
      else revert();                                              
    }                                                          
    p = offset;                                                 
    r.v = new uint8[](counters[1]);
    r.r = new bytes32[](counters[2]);
    r.s = new bytes32[](counters[3]);
                                                    
    while(p < offset+sz) {                                     
      (fieldId, wireType, bytesRead) = _pb._decode_key(p, bs);  
      p += bytesRead;                                           
      if (false) {}
      else if(fieldId == 1)       
          p += _read_v(p, bs, r, counters);
      else if(fieldId == 2)       
          p += _read_r(p, bs, r, counters);
      else if(fieldId == 3)       
          p += _read_s(p, bs, r, counters);
      else revert();                                             
    }                                                          
    return (r, sz);                                             
  }                                                            
                            
  // field readers                       
  function _read_v(uint p, bytes bs, Data r, uint[4] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_uint8(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[1] += 1;                                            
    } else {                                                         
      r.v[ r.v.length - counters[1] ] = x;                                         
      if(counters[1] > 0) counters[1] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_r(uint p, bytes bs, Data r, uint[4] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_bytes32(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[2] += 1;                                            
    } else {                                                         
      r.r[ r.r.length - counters[2] ] = x;                                         
      if(counters[2] > 0) counters[2] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
  function _read_s(uint p, bytes bs, Data r, uint[4] counters) internal constant returns (uint) {                            
    var (x, sz) = _pb._decode_sol_bytes32(p, bs);                                  
    if(isNil(r)) {                                                  
      counters[3] += 1;                                            
    } else {                                                         
      r.s[ r.s.length - counters[3] ] = x;                                         
      if(counters[3] > 0) counters[3] -= 1;                      
    }                                                                
    return sz;                                                       
  }                                                                 
                            
  // struct decoder                       
                                      
            
  //store function                                                     
  function store(Data memory input, Data storage output) internal{
    output.v = input.v;                           
    output.r = input.r;                           
    output.s = input.s;                           
  }                                                                   
             
  //utility functions                                           
  function nil() internal constant returns (Data r) {        
    assembly { r := 0 }                                       
  }                                                            
  function isNil(Data x) internal constant returns (bool r) {
    assembly { r := iszero(x) }                               
  }                                                            
} 
