package bitcoin

//BlockScanNotificationObject 扫描被通知对象
type BTCBlockScanNotificationObject interface {

	//BlockScanNotify 新区块扫描完成通知
	//@required
	BTCBlockScanNotify(block *Block, txs []*Transaction) error
}

////ExtractTxOriginResult
//type ExtractTxOriginResult struct {
//	Tx      *Transaction
//	Success bool
//}
//
////SetScanBlockTaskOrigin 扫描任务
//func (bs *BTCBlockScanner) SetScanBlockTaskOrigin() {
//	bs.SetTask(bs.ScanBlockTaskOrigin)
//}
//
////ScanBlockTask 扫描任务
//func (bs *BTCBlockScanner) ScanBlockTaskOrigin() {
//
//	//获取本地区块高度
//	blockHeader, err := bs.GetScannedBlockHeaderOrigin()
//	if err != nil {
//		bs.wm.Log.Std.Info("block scanner can not get new block height; unexpected error: %v", err)
//		return
//	}
//
//	currentHeight := blockHeader.Height
//	currentHash := blockHeader.Hash
//
//	for {
//
//		if !bs.Scanning {
//			//区块扫描器已暂停，马上结束本次任务
//			return
//		}
//
//		//获取最大高度
//		maxHeight, err := bs.wm.GetBlockHeight()
//		if err != nil {
//			//下一个高度找不到会报异常
//			bs.wm.Log.Std.Info("block scanner can not get rpc-server block height; unexpected error: %v", err)
//			break
//		}
//
//		//是否已到最新高度
//		if currentHeight >= maxHeight {
//			bs.wm.Log.Std.Info("block scanner has scanned full chain data. Current height: %d", maxHeight)
//			break
//		}
//
//		//继续扫描下一个区块
//		currentHeight = currentHeight + 1
//
//		bs.wm.Log.Std.Info("block scanner scanning height: %d ...", currentHeight)
//
//		hash, err := bs.wm.GetBlockHash(currentHeight)
//		if err != nil {
//			//下一个高度找不到会报异常
//			bs.wm.Log.Std.Info("block scanner can not get new block hash; unexpected error: %v", err)
//			break
//		}
//
//		block, err := bs.wm.getBlockByCore(hash, 2)
//		if err != nil {
//			bs.wm.Log.Std.Info("block scanner can not get new block data; unexpected error: %v", err)
//			break
//		}
//
//		isFork := false
//
//		//判断hash是否上一区块的hash
//		if currentHash != block.Previousblockhash && len(currentHash) > 0 {
//
//			bs.wm.Log.Std.Info("block has been fork on height: %d.", currentHeight)
//			bs.wm.Log.Std.Info("block height: %d local hash = %s ", currentHeight-1, currentHash)
//			bs.wm.Log.Std.Info("block height: %d mainnet hash = %s ", currentHeight-1, block.Previousblockhash)
//
//			bs.wm.Log.Std.Info("delete recharge records on block height: %d.", currentHeight-1)
//
//			//查询本地分叉的区块
//			forkBlock, _ := bs.wm.GetLocalBlock(currentHeight - 1)
//
//			currentHeight = currentHeight - 2 //倒退2个区块重新扫描
//			if currentHeight <= 0 {
//				currentHeight = 1
//			}
//
//			localBlock, err := bs.wm.GetLocalBlock(currentHeight)
//			if err != nil {
//				bs.wm.Log.Std.Error("block scanner can not get local block; unexpected error: %v", err)
//
//				//查找core钱包的RPC
//				bs.wm.Log.Info("block scanner prev block height:", currentHeight)
//
//				prevHash, err := bs.wm.GetBlockHash(currentHeight)
//				if err != nil {
//					bs.wm.Log.Std.Error("block scanner can not get prev block; unexpected error: %v", err)
//					break
//				}
//
//				localBlock, err = bs.wm.getBlockByCore(prevHash, 1)
//				if err != nil {
//					bs.wm.Log.Std.Error("block scanner can not get prev block; unexpected error: %v", err)
//					break
//				}
//
//			}
//
//			//重置当前区块的hash
//			currentHash = localBlock.Hash
//
//			bs.wm.Log.Std.Info("rescan block on height: %d, hash: %s .", currentHeight, currentHash)
//
//			//重新记录一个新扫描起点
//			bs.wm.SaveLocalNewBlock(localBlock.Height, localBlock.Hash)
//
//			isFork = true
//
//			if forkBlock != nil {
//
//				//通知分叉区块给观测者，异步处理
//				bs.NewBTCBlockNotify(forkBlock, isFork)
//			}
//
//		} else {
//
//			if !block.isVerbose {
//				err = bs.BatchExtractTransactionOrigin(block)
//				if err != nil {
//					bs.wm.Log.Std.Info("block scanner can not extractRechargeRecords; unexpected error: %v", err)
//				}
//			}
//
//			//重置当前区块的hash
//			currentHash = hash
//
//			//保存本地新高度
//			bs.wm.SaveLocalNewBlock(currentHeight, currentHash)
//			bs.wm.SaveLocalBlock(block)
//
//			isFork = false
//
//			//通知新区块给观测者，异步处理
//			bs.NewBTCBlockNotify(block, isFork)
//		}
//
//	}
//
//	if bs.IsScanMemPool {
//		//扫描交易内存池
//		bs.ScanTxMemPoolOrigin()
//	}
//}
//
////ScanTxMemPool 扫描交易内存池
//func (bs *BTCBlockScanner) ScanTxMemPoolOrigin() {
//
//	bs.wm.Log.Std.Info("block scanner scanning mempool ...")
//
//	//提取未确认的交易单
//	txIDsInMemPool, err := bs.wm.GetTxIDsInMemPool()
//	if err != nil {
//		bs.wm.Log.Std.Info("block scanner can not get mempool data; unexpected error: %v", err)
//		return
//	}
//
//	if txIDsInMemPool == nil || len(txIDsInMemPool) == 0 {
//		return
//	}
//
//	block := &Block{
//		Height:    0,
//		Hash:      "",
//		isVerbose: false,
//		tx:        txIDsInMemPool,
//	}
//
//	err = bs.BatchExtractTransactionOrigin(block)
//	if err != nil {
//		bs.wm.Log.Std.Info("block scanner can not extractRechargeRecords; unexpected error: %v", err)
//	}
//
//	//通知内存池交易单给观测者
//	bs.NewBTCBlockNotify(block, false)
//}
//
////BatchExtractTransaction 批量提取交易单
////bitcoin 1M的区块链可以容纳3000笔交易，批量多线程处理，速度更快
//func (bs *BTCBlockScanner) BatchExtractTransactionOrigin(block *Block) error {
//
//	var (
//		done       = 0 //完成标记
//		shouldDone = len(block.tx) //需要完成的总数
//	)
//
//	if len(block.tx) == 0 {
//		return fmt.Errorf("BatchExtractTransaction block is nil.")
//	}
//
//	//生产通道
//	producer := make(chan interface{})
//	//defer close(producer)
//
//	//消费通道
//	worker := make(chan interface{})
//	defer close(worker)
//
//	//保存工作
//	saveWork := func(result chan interface{}) {
//
//		for {
//			select {
//			case obj, exist := <-result:
//				if !exist {
//					return
//				}
//				txResult, ok := obj.(*ExtractTxOriginResult)
//				if ok {
//
//					if txResult.Success {
//						block.txDetails = append(block.txDetails, txResult.Tx)
//						//bs.wm.Log.Debugf("txDetails Length = %d", len(block.txDetails))
//					}
//
//				}
//				//累计完成的线程数
//				done++
//				//bs.wm.Log.Std.Info("done = %d, shouldDone = %d ", done, shouldDone)
//				if done == shouldDone {
//
//					close(producer) //关闭通道，等于给通道传入nil
//				}
//			}
//
//		}
//	}
//
//	//提取工作
//	extractWork := func(mTxs []string, eProducer chan interface{}) {
//		for _, txid := range mTxs {
//			bs.extractingCH <- struct{}{}
//			//shouldDone++
//			go func(mTxid string, end chan struct{}, mProducer chan<- interface{}) {
//
//				//释放
//				defer func() {
//					<-end
//				}()
//
//				result := &ExtractTxOriginResult{
//					Success: true,
//				}
//
//				//导出提出的交易
//				trx, err := bs.wm.GetTransaction(txid)
//				if err != nil {
//					bs.wm.Log.Std.Info("block scanner get transaction txid: %s failed, err: %v", txid, err)
//					result.Success = false
//				}
//				result.Tx = trx
//
//				mProducer <- result
//
//			}(txid, bs.extractingCH, eProducer)
//		}
//	}
//
//	/*	开启导出的线程	*/
//
//	//独立线程运行消费
//	go saveWork(worker)
//
//	//独立线程运行生产
//	go extractWork(block.tx, producer)
//
//	//以下使用生产消费模式
//	concurrent.ProducerToConsumerRuntime(producer, worker)
//
//	return nil
//}
//
////AddObserver 添加观测者
//func (bs *BTCBlockScanner) AddBTCBlockObserver(obj BTCBlockScanNotificationObject) error {
//	bs.Mu.Lock()
//
//	defer bs.Mu.Unlock()
//
//	if obj == nil {
//		return nil
//	}
//	if _, exist := bs.BTCBlockObservers[obj]; exist {
//		//已存在，不重复订阅
//		return nil
//	}
//
//	bs.BTCBlockObservers[obj] = true
//
//	return nil
//}
//
////RemoveObserver 移除观测者
//func (bs *BTCBlockScanner) RemoveBTCBlockObserver(obj BTCBlockScanNotificationObject) error {
//	bs.Mu.Lock()
//	defer bs.Mu.Unlock()
//
//	delete(bs.BTCBlockObservers, obj)
//
//	return nil
//}
//
////newBlockNotify 获得新区块后，通知给观测者
//func (bs *BTCBlockScanner) NewBTCBlockNotify(block *Block, isFork bool) {
//	block.Fork = isFork
//	for o, _ := range bs.BTCBlockObservers {
//		o.BTCBlockScanNotify(block, block.txDetails)
//	}
//}
//
//
////GetScannedBlockHeader 获取当前扫描的区块头
//func (bs *BTCBlockScanner) GetScannedBlockHeaderOrigin() (*openwallet.BlockHeader, error) {
//
//	var (
//		blockHeight uint64 = 0
//		hash        string
//	)
//
//	blockHeight, hash = bs.wm.GetLocalNewBlock()
//
//	return &openwallet.BlockHeader{Height: blockHeight, Hash: hash}, nil
//}
