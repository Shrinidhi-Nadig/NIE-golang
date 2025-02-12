async function checkexamresults() {
    return new Promise((resolve,reject)=>{
        setTimeout(()=>{
            const passed=Math.random()>0.35;
            if(passed){
                resolve("Congratulations!! You have passed the exam..");
            }else{
                reject("Better Luck next time!, you've Failed")
            }
        },2000);
    })
}

async function getexamresults() {
    try{
        const result=await checkexamresults();
        console.log(result)
    }catch(error){
        console.error(error);
    }
}

getexamresults();