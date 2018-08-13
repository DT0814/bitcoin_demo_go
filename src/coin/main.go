package main

import "core"

/*func main1() {
	bc := core.NewBlockChain()

	bc.AddBlock("one")
	bc.AddBlock("two")

	for _, block := range bc.Blocks {
		fmt.Printf("prevHash %x \n", block.PrevBlockHash)
		fmt.Printf("data     %s \n", block.Data)
		fmt.Printf("Hash     %x \n", block.Hash)
		fmt.Println()
	}
}*/
/*func main2() {
	unix1 := time.Now().Unix()
	fmt.Printf("%d\n", unix1)
	bc := core.NewBlockChain()
	unix1stop := time.Now().Unix()
	fmt.Printf("%d\n", unix1stop)
	bc.AddBlock("one")
	unix2Stop := time.Now().Unix()
	fmt.Printf("%d\n", unix2Stop)
	bc.AddBlock("two")
	unix3Stop := time.Now().Unix()
	fmt.Printf("%d \n", unix3Stop)
	for _, block := range bc.Blocks {
		fmt.Printf("prevHash %x \n", block.PrevBlockHash)
		fmt.Printf("data     %s \n", block.Data)
		fmt.Printf("Hash     %x \n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("pow : %s \n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}

}
*/
func main() {
	bc := core.NewBlockChain()
	defer bc.Db.Close()

	cli := core.CLI{bc}
	cli.Run()

}
