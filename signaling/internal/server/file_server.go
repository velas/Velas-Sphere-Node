package server

import (
	"fmt"
	"signaling/internal/codec"
	"signaling/internal/entity"

	"github.com/pion/webrtc/v3"
)

func FileServer() {

	const messageSize = 15

	api := webrtc.NewAPI()

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		fmt.Println("failed to create the peer connection", err)
		return
	}

	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		fmt.Println("failed to create the data channel")
		return
	}

	peerConnection.OnICEConnectionStateChange(
		func(connectionState webrtc.ICEConnectionState) {
			fmt.Printf("ICE Connection State has changed: %s\n", connectionState.String())
		},
	)

	dataChannel.OnOpen(
		func() {
			fmt.Printf("Data channel '%s'-'%d' open.\n", dataChannel.Label(), dataChannel.ID())
		},
	)

	dataChannel.OnMessage(
		func(msg webrtc.DataChannelMessage) {
			message, err := codec.DecodeMessage(msg.Data)
			if err != nil {
				fmt.Println("failed to decode the message:", err)
				return
			}

			switch body := message.Body.(type) {
			case entity.BlockTransmissionMessage:
				dataLen := len(body.Data)
				fmt.Println("got data with len", dataLen)

			case entity.UploadInitializationMessage:
				hash := len(body.Hash)
				fmt.Println("got data with hash", hash)

				// TODO: send data
				// dataChannel.Send()

			case entity.DownloadInitializationMessage:
				hash := len(body.Hash)
				fmt.Println("got data with hash", hash)
			default:
				fmt.Println("unknown message")
			}
		},
	)

	// peerConnection.OnICECandidate(
	// 	func(c *webrtc.ICECandidate) {
	// 		fmt.Println(c)
	// 	},
	// )

	// Create channel that is blocked until ICE Gathering is complete
	gatherComplete := webrtc.GatheringCompletePromise(peerConnection)

	// TODO: change the flow?

	// offer collection
	go func() {
		OfferServer(peerConnection)
	}()
	// Block until ICE Gathering is complete, disabling trickle ICE
	// we do this because we only can exchange one signaling message
	// in a production application you should exchange ICE Candidates via OnICECandidate
	<-gatherComplete

	// Block forever
	select {}
}
