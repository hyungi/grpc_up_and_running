# How to build & run
It is basically same with [chapter 2](../ch02/RUNBOOK.md)

## Server side message
```shell
2021/12/20 22:17:06 Order Added. ID : 101
2021/12/20 22:17:06 key: 106, order: id:"106" items:"Amazon Echo" items:"Apple iPhone XS" price:300 destination:"Mountain View, CA"
2021/12/20 22:17:06 key: 107, order: id:"107" items:"Lenovo ThinkPad" items:"Galaxy S21" price:1700 destination:"Mountain View, CA"
2021/12/20 22:17:06 key: 108, order: id:"108" items:"Dell XPS" price:1000 destination:"San Diego, CA"
2021/12/20 22:17:06 key: 101, order: id:"101" items:"iPhone XS" items:"Mac Book Pro" price:2300 destination:"San Jose, CA"
2021/12/20 22:17:06 key: 102, order: id:"102" items:"Google Pixel 3A" items:"Mac Book Pro" price:1800 destination:"Mountain View, CA"
2021/12/20 22:17:06 Matching Order Found : 102
2021/12/20 22:17:06 key: 103, order: id:"103" items:"Apple Watch S4" price:400 destination:"San Jose, CA"
2021/12/20 22:17:06 key: 104, order: id:"104" items:"Google Home Mini" items:"Google Nest Hub" price:400 destination:"Mountain View, CA"
2021/12/20 22:17:06 Matching Order Found : 104
2021/12/20 22:17:06 key: 105, order: id:"105" items:"Amazon Echo" price:30 destination:"San Jose, CA"
2021/12/20 22:17:06 Order ID : 102 - Updated
2021/12/20 22:17:06 Order ID : 103 - Updated
2021/12/20 22:17:06 Order ID : 104 - Updated
2021/12/20 22:17:06 Reading Proc order : value:"102"
2021/12/20 22:17:06 1cmb - Mountain View, CA
2021/12/20 22:17:06 Reading Proc order : value:"103"
2021/12/20 22:17:06 1cmb - San Jose, CA
2021/12/20 22:17:06 Reading Proc order : value:"104"
2021/12/20 22:17:06 Shipping : cmb - Mountain View, CA -> 2
2021/12/20 22:17:06 Shipping : cmb - San Jose, CA -> 1
2021/12/20 22:17:07 Reading Proc order : value:"101"
2021/12/20 22:17:07 1cmb - San Jose, CA
2021/12/20 22:17:07 Reading Proc order : value:"106"
2021/12/20 22:17:07 1cmb - Mountain View, CA
2021/12/20 22:17:07 Reading Proc order : value:"105"
2021/12/20 22:17:07 Shipping : cmb - San Jose, CA -> 2
2021/12/20 22:17:07 Shipping : cmb - Mountain View, CA -> 1
2021/12/20 22:17:07 Reading Proc order : value:"107"
2021/12/20 22:17:07 1cmb - Mountain View, CA
2021/12/20 22:17:07 Reading Proc order : value:"108"
2021/12/20 22:17:07 1cmb - San Diego, CA
2021/12/20 22:17:07 Reading Proc order : <nil>
2021/12/20 22:17:07 EOF : <nil>
```

## Client side message
```shell
2021/12/20 21:55:57 ===== Start Simple RPC test =====                                                                                        
2021/12/20 21:55:57 AddOrder Response -> Order Added: 101                                                                                    
2021/12/20 21:55:57 GetOrder Response -> : id:"106"  items:"Amazon Echo"  items:"Apple iPhone XS"  price:300  destination:"Mountain View, CA"
2021/12/20 21:55:57 ===== Start Server Streaming RPC test  =====                                                                             
2021/12/20 21:55:57 Search Result : id:"104"  items:"Google Home Mini"  items:"Google Nest Hub"  price:400  destination:"Mountain View, CA"                                                                                                                                         
2021/12/20 21:55:57 Search Result : id:"102"  items:"Google Pixel 3A"  items:"Mac Book Pro"  price:1800  destination:"Mountain View, CA"                                                                                                                                            
2021/12/20 21:55:57 EOF                                                                                                                                                                                                                                                             
2021/12/20 21:55:57 ===== Start Client Streaming RPC test  =====                                                                                                                                                                                                                    
2021/12/20 21:55:57 Update Orders Res : value:"Orders processed Updated Order IDs : 102, 103, 104"                                                                                                                                                                                  
2021/12/20 21:55:57 ===== Start Bi-di Streaming RPC test  =====                                                                                                                                                                                                                     
2021/12/20 21:55:57 Combined shipment : [id:"102"  items:"Google Pixel 3A"  items:"Google Pixel Book"  price:1100  destination:"Mountain View, CA",     id:"104"  items:"Google Home Mini"  items:"Google Nest Hub"  items:"iPad Mini"  price:2200  destination:"Mountain View, CA"]
2021/12/20 21:55:57 Combined shipment : [id:"103"  items:"Apple Watch S4"  items:"Mac Book Pro"  items:"iPad Pro"  price:2800  destination:"San Jose, CA"]                                                                                                                          
2021/12/20 21:55:58 Combined shipment : [id:"101"  items:"iPhone XS"  items:"Mac Book Pro"  price:2300  destination:"San Jose, CA",     id:"105"  items:"Amazon Echo"  price:30  destination:"San Jose, CA"]
2021/12/20 21:55:58 Combined shipment : [id:"106"  items:"Amazon Echo"  items:"Apple iPhone XS"  price:300  destination:"Mountain View, CA"]                                                                
2021/12/20 21:55:58 Combined shipment : [id:"108"  items:"Dell XPS"  price:1000  destination:"San Diego, CA"]
2021/12/20 21:55:58 Combined shipment : [id:"107"  items:"Lenovo ThinkPad"  items:"Galaxy S21"  price:1700  destination:"Mountain View, CA"]
```