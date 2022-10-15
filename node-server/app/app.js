const express = require("express");
const cors = require("cors");
const app = express();
const {
    TodoCollection
} = require("../db/index.js");
app.use(cors());
app.use(express.json());

app.get("/", async (req, res) => {
    const data = await TodoCollection.find({});
    console.log(data);
    return res.status(200).json(data);
});

/* app.put("/:id", async (req, res) => { */
/*     try { */
/*         let {id} = req.url */
/*         return res.status(200).json(true); */
/*     } catch (error) { */
/*         console.log(error); */
/*         return res.status(400).json(false); */
/*     } */
/* }); */

app.post("/", async (req, res) => {
    try {
        let done = await TodoCollection.insertMany(req.body);
        console.log(done);
        return res.status(200).json({
            done,
        });
    } catch (error) {
        console.log(error);
        return res.status(400).json(false);
    }
});

app.delete("/:id", async (req, res) => {
    try {
        let {id} = req.url
        let done = await TodoCollection.findByIdAndDelete({
            id
        });
        return res.status(200).json(true);
    } catch (error) {
        console.log(error);
        return res.status(400).json(false);
    }
});

app.delete("/", async (req, res) => {
    try {
        let done = await TodoCollection.deleteMany({});
        return res.status(200).json(true);
    } catch (error) {
        console.log(error);
        return res.status(400).json(false);
    }
});

module.exports = {
    app,
};
