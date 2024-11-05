package main

import (
	"context"
	"go-cms-gql/database"
	"go-cms-gql/directives"
	"go-cms-gql/graph/model"
	"go-cms-gql/utils"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const userCollection = utils.USER_COLLECTION

type adminRequest struct {
	Username string `validate:"required,min=3,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,containsNumber,containsSpecialCharacter"`
}

func main() {
	directives.InitValidator()

	ctx := context.TODO()
	input := model.NewUser{
		Username: utils.GetValue("ADMIN_NAME"),
		Email:    utils.GetValue("ADMIN_EMAIL"),
		Password: utils.GetValue("ADMIN_PASSWORD"),
	}

	if os.Getenv("APP_MODE") != "production" {
		createAdmin(ctx, input)
	} else {
		execInitAdmin(input)
	}
}

func connectToDB() {
	err := database.Connect(utils.GetValue("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v\n", err)
	}
}

func createAdmin(ctx context.Context, input model.NewUser) {
	connectToDB()

	if err := directives.ValidateStruct(&adminRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		log.Fatalf("validation failed: %v\n", err)
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error occurred when creating password: %v\n", err)
	}

	var password string = string(bs)

	var newUser model.User = model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  password,
		Role:      utils.ADMIN_ROLE,
		CreatedAt: time.Now(),
	}

	var collection *mongo.Collection = database.GetCollection(userCollection)

	var foundUser *model.User = &model.User{}
	userFilter := bson.M{"role": utils.ADMIN_ROLE}

	err = collection.FindOne(ctx, userFilter).Decode(foundUser)

	if err == nil {
		log.Fatal("admin already exists")
	} else if err != mongo.ErrNoDocuments {
		log.Fatalf("error occurred when fetching document: %v\n", err)
	}

	res, err := collection.InsertOne(ctx, newUser)

	if err != nil || res.InsertedID == nil {
		log.Fatalf("admin creation failed: %v\n", err)
	}

	log.Println("admin created successfully")
}

func execInitAdmin(input model.NewUser) {
	bs, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error occurred when creating password: %v\n", err)
	}

	var password string = string(bs)

	mongoRootUsername := utils.GetValue("MONGO_INITDB_ROOT_USERNAME")
	mongoRootPassword := utils.GetValue("MONGO_INITDB_ROOT_PASSWORD")
	adminName := input.Username
	adminEmail := input.Email
	adminPassword := password

	// Determine the appropriate command to execute the script based on OS
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// On Windows, use "bash" to execute the shell script
		cmd = exec.Command("bash", "./init_admin.sh", mongoRootUsername, mongoRootPassword, adminName, adminEmail, adminPassword)
	} else {
		// On Unix-based systems, execute the init_admin directly
		cmd = exec.Command("./init_admin.sh", mongoRootUsername, mongoRootPassword, adminName, adminEmail, adminPassword)
	}

	// Set the environment variables if needed
	cmd.Env = os.Environ()

	// Capture output and errors
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Script execution failed: %v\nOutput: %s", err, output)
	}

	// Print the output from the script
	log.Printf("Script executed successfully:\n%s", output)

	log.Println("admin created successfully")
}
