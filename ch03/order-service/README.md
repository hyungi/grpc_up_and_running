# There is an error in original code.

## On server

```golang
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	...
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order : %s", orderId)
		
		// handling orders

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}
```

As you can see, original code initialize `batchMarker` as 0 when Batch job is done.
But, it causes malfunction batch job.

Below is result of original `ProcessOrders` function result.
And server put order id 106 and 107 in same shipment.
But, the order of orders is `[102, 103, 104]`, `[101,106,105]`, `[107,108]`
According to author's definition, `orderBatchSize` is not same with size of `Combined shipment`.
```shell
2021/12/20 21:28:01 ===== Start Bi-di Streaming RPC test  =====
2021/12/20 21:28:01 Combined shipment : [id:"102"  items:"Google Pixel 3A"  items:"Google Pixel Book"  price:1100  destination:"Mountain View, CA",     id:"104"  items:"Google Home Mini"  items:"Google Nest Hub"  items:"iPad Mini"  price:2200  destination:"Mountain View, CA"]
2021/12/20 21:28:01 Combined shipment : [id:"103"  items:"Apple Watch S4"  items:"Mac Book Pro"  items:"iPad Pro"  price:2800  destination:"San Jose, CA"]
2021/12/20 21:28:02 Combined shipment : [id:"101"  items:"iPhone XS"  items:"Mac Book Pro"  price:2300  destination:"San Jose, CA",     id:"105"  items:"Amazon Echo"  price:30  destination:"San Jose, CA"]
2021/12/20 21:28:02 Combined shipment : [id:"106"  items:"Amazon Echo"  items:"Apple iPhone XS"  price:300  destination:"Mountain View, CA",    id:"107"  items:"Lenovo ThinkPad"  items:"Galaxy S21"  price:1700  destination:"Mountain View, CA"]
2021/12/20 21:28:02 Combined shipment : [id:"108"  items:"Dell XPS"  price:1000  destination:"San Diego, CA"]
```

So, it should be like this.
```shell
2021/12/20 21:21:52 Combined shipment : [id:"102"  items:"Google Pixel 3A"  items:"Google Pixel Book"  price:1100  destination:"Mountain View, CA",     id:"104"  items:"Google Home Mini"  items:"Google Nest Hub"  items:"iPad Mini"  price:2200  destination:"Mountain View, CA"]
2021/12/20 21:21:52 Combined shipment : [id:"103"  items:"Apple Watch S4"  items:"Mac Book Pro"  items:"iPad Pro"  price:2800  destination:"San Jose, CA"]
2021/12/20 21:21:53 Combined shipment : [id:"101"  items:"iPhone XS"  items:"Mac Book Pro"  price:2300  destination:"San Jose, CA",     id:"105"  items:"Amazon Echo"  price:30  destination:"San Jose, CA"]
2021/12/20 21:21:53 Combined shipment : [id:"106"  items:"Amazon Echo"  items:"Apple iPhone XS"  price:300  destination:"Mountain View, CA"]
2021/12/20 21:21:53 Combined shipment : [id:"107"  items:"Lenovo ThinkPad"  items:"Galaxy S21"  price:1700  destination:"Mountain View, CA"]
2021/12/20 21:21:53 Combined shipment : [id:"108"  items:"Dell XPS"  price:1000  destination:"San Diego, CA"]
```

So, I'm gonna re-initialize batchMarker as 1 not 0