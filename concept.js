class Game{
    constructor(player1, player2){
        this.player1 = player1;
        this.player2 = player2;
        this.gameIsActive = true;
        this.winnerId = 0; 
    }

    endGame(winnerId){
        this.gameIsActive = false;
        this.winnerId = winnerId;
        //somehow end game
    }
}

class Card{
    constructor(cost, damage, hp){
        this.cost = cost
        this.damage = damage 
        this.hp = hp
    }
}

class Player{
    constructor(){
        this.hp = 30;
        this.deck = [];
        this.hand = [];
        for (let i = 1; i < 11; i++) this.deck.push(new Card(i,i,i));
    }

    DrawCards(num) {
        for (let j = 0; j < num; j++)if (this.deck.length > 0) this.hand.push(this.deck.pop());
    }
}

let player1 = new Player();
let player2 = new Player();
let game = new Game(player1, player2);

while (true){//game loop
    for (let i=1;game.gameIsActive;i++){//turns
        let mana = (i<10)?i:10;
        player1.Mana = mana;//set player mana to turn's val
        player1.DrawCards(1);
        while (true){//player1's turn
            userInput = scanForUserInput();//scan for user input
            if (userInput.valid()){//check if input given is valid(card is playable, board has enough space, card has target, etc.)
                userInput.selectedCard.play();//play card selected, remove mana etc.
            }else if (userInput.isTurnEnd()){
                break; //end player1's turn
            }
        }
        
        player2.Mana = mana;//set player mana to turn's val
        player2.DrawCards(1);
        while (true){//player2's turn
            userInput = scanForUserInput();//scan for user input
            if (userInput.valid()){//check if input given is valid(card is playable, board has enough space, card has target, etc.)
                userInput.selectedCard.play();//play card selected, remove mana etc.
            }else if (userInput.isTurnEnd()){
                break; //end player2's turn
            }
        }
    }
}

