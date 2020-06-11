package ffirebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/pkg/errors"
	"github.com/reynaldoqs/urLix_resolver/internal/core/domain"
	"google.golang.org/api/option"
)

type firebaseService struct {
	store    *firestore.Client
	cloudMsg *messaging.Client
}

func NewFirebaseApp(configPath string) *firebaseService {
	cOptions := option.WithCredentialsFile(configPath)

	firebaseApp, err := firebase.NewApp(context.Background(), nil, cOptions)
	if err != nil {
		err = errors.Wrap(err, "firebase.NewFirebaseApp")
		log.Fatalln(err)
	}

	firestoreCl, err := firebaseApp.Firestore(context.Background())
	if err != nil {
		err = errors.Wrap(err, "firebase.NewFirebaseApp")
		log.Fatalln(err)
	}

	messagingCl, err := firebaseApp.Messaging(context.Background())
	if err != nil {
		err = errors.Wrap(err, "firebase.NewFirebaseApp")
		log.Fatalln(err)
	}

	return &firebaseService{
		store:    firestoreCl,
		cloudMsg: messagingCl,
	}
}

// repository implementations

func (fs *firebaseService) GetAll() ([]*domain.Farmer, error) {

	collection := fs.store.Collection("farmerResolvers")

	var farmers []*domain.Farmer

	result, err := collection.Documents(context.TODO()).GetAll()
	if err != nil {
		err = errors.Wrap(err, "firebase.GetAll")
		return nil, err
	}

	for _, docSnapshot := range result {
		var farmer domain.Farmer
		docSnapshot.DataTo(&farmer)
		farmers = append(farmers, &farmer)
	}

	return farmers, nil
}

func (fs *firebaseService) GetByNumber(num int) (*domain.Farmer, error) {
	collection := fs.store.Collection("farmerResolvers")

	var farmer domain.Farmer

	dsnap, err := collection.Doc(strconv.Itoa(num)).Get(context.TODO())
	if err != nil {
		err = errors.Wrap(err, "firebase.GetByNumber")
		return nil, err
	}

	err = dsnap.DataTo(&farmer)

	if err != nil {
		err = errors.Wrap(err, "firebase.GetByNumber")
		return nil, err
	}

	return &farmer, nil
}

// Cloud messagin implementations
type RechargeMsgStandart struct {
	ExecCodes     string `json:"execCodes"`
	TargetCompany string `json:"targetCompany"`
	IDRecharge    string `json:"idRecharge"`
	Mount         string `json:"mount"`
	FarmerNumber  string `json:"farmerNumber"`
}

func (fs *firebaseService) RechargeNotify(farmer *domain.Farmer, fcmsg *domain.RechargeMessage) error {
	fmt.Println("staring")
	notification := messaging.Notification{
		Title: farmer.DeviceID,
		Body:  fmt.Sprintf("el numero %v necesita recarga", farmer.PhoneNumber),
	}

	temp := RechargeMsgStandart{
		ExecCodes:     strings.Join(fcmsg.ExecCodes, "&"),
		TargetCompany: fcmsg.TargetCompany,
		IDRecharge:    fcmsg.IDRecharge,
		Mount:         strconv.Itoa(fcmsg.Mount),
		FarmerNumber:  strconv.Itoa(fcmsg.FarmerNumber),
	}

	out, err := json.Marshal(temp)
	if err != nil {
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	fmt.Println(string(out))

	m := make(map[string]string)

	err = json.Unmarshal(out, &m)
	if err != nil {
		fmt.Println(err)
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	message := messaging.Message{
		Token:        farmer.MsgToken,
		Data:         m,
		Notification: &notification,
	}

	_, err = fs.cloudMsg.Send(context.TODO(), &message)
	if err != nil {
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	return err
}

type AdminMsgStandart struct {
	ExecCodes    string `json:"execCodes"`
	IDMessage    string `json:"idMessage"`
	FarmerNumber string `json:"farmerNumber"`
}

func (fs *firebaseService) AdminNotify(farmer *domain.Farmer, amsg *domain.AdminMessage) error {
	fmt.Println("staring")
	notification := messaging.Notification{
		Title: farmer.DeviceID,
		Body:  fmt.Sprintf("el numero %v necesita recarga", farmer.PhoneNumber),
	}

	temp := AdminMsgStandart{
		ExecCodes:    strings.Join(amsg.ExecCodes, "&"),
		IDMessage:    amsg.IDMessage,
		FarmerNumber: strconv.Itoa(amsg.FarmerNumber),
	}

	out, err := json.Marshal(temp)
	if err != nil {
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	fmt.Println(string(out))

	m := make(map[string]string)

	err = json.Unmarshal(out, &m)
	if err != nil {
		fmt.Println(err)
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	message := messaging.Message{
		Token:        farmer.MsgToken,
		Data:         m,
		Notification: &notification,
	}

	_, err = fs.cloudMsg.Send(context.TODO(), &message)
	if err != nil {
		err = errors.Wrap(err, "firebase.Notify")
		return err
	}

	return err
}
