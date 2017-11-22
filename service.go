package capusta

type Iterator struct {
	blockchain *blockchain
	curBlock   *Block
	curIndex   int
	lastIndex  int
}

func NewItrator(bc *blockchain) *Iterator {

	iterator := Iterator{}

	iterator.blockchain = bc
	iterator.lastIndex = bc.getLenght() - 1
	iterator.curIndex = 0
	iterator.curBlock = bc.getBlockbyIndex(iterator.curIndex)

	return &iterator
}

func (iter *Iterator) HasNext() bool {
	return iter.curIndex <= iter.lastIndex
}

func (iter *Iterator) Next() *Block {
	iter.curIndex += 1
	iter.curBlock = iter.blockchain.getBlockbyIndex(iter.curIndex)
	return iter.curBlock
}

func (iter *Iterator) GetBlock() *Block {
	return iter.curBlock
}
