package controllers

import (
	"zd112_api/models"
	"time"
	"github.com/davecgh/go-spew/spew"
	"github.com/astaxie/beego"
)

type BlockChainController struct {
	BaseController
}

var BlockChain []models.Block

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :userId is empty
// @router / [get]
func (this *BlockChainController) Get() {
	//this.checkToken()
	beego.Info("------------------block~get")
	genesisBlock := models.Block{0, time.Now().Unix(), 0, "", ""}
	spew.Dump(genesisBlock)
	BlockChain = append(BlockChain, genesisBlock)
	this.response(genesisBlock)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *BlockChainController) Post() {
	bpm := this.getInt("bpm", 0)
	if bpm > 0 {
		block := new(models.Block)
		if newBlock, err := block.GenerateBlock(BlockChain[len(BlockChain)-1], bpm); err != nil {
			this.false(-1, err)
		} else {
			if block.IsBlockValid(newBlock, BlockChain[len(BlockChain)-1]) {
				newBlockChain := append(BlockChain, newBlock)
				this.replaceChain(newBlockChain)
				spew.Dump(BlockChain)
				this.response(newBlock)
			} else {
				this.false(-1, "this is block valid false!")
			}
		}
	} else {
		this.false(-1, bpm)
	}
}

func (this *BlockChainController) replaceChain(newBlocks []models.Block) {
	if len(newBlocks) > len(BlockChain) {
		BlockChain = newBlocks
	}
}
