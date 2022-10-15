const { Schema, connect, model } = require("mongoose");

connect("mongodb://localhost:27017/VueToDo")
  .then(() => {
        console.log('connected')
    })
  .catch((err) => {
        console.log(err)
    });
const toDoSchema = new Schema({
    act:{
        type:String,
        required:true
    },
    done:{
        type:Boolean,
        default:false
    }
}, {
    timestamps: true
});

const TodoCollection=model('VueToDo',toDoSchema)

module.exports={
    TodoCollection
}
