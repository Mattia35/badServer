<script>
export default{
    data: function(){
        return{
            errormsg: null,
        };
    },
    methods: {
        async goToEmployees(){
            this.errormsg = null;
            try {
                // redirect to the employees page
                this.$router.push(`/${sessionStorage.username}/employees`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
        async goToProjects(){
            this.errormsg = null;
            try {
                // redirect to the projects page
                this.$router.push(`/${sessionStorage.username}/projects`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
        async goToDepartments(){
            this.errormsg = null;
            try {
                // redirect to the departments page
                this.$router.push(`/${sessionStorage.username}/departments`);
            } catch (e) {
                this.errormsg = e.toString();
            };
        },
    }
}
</script>

<template>
    <div class="container">
      <div v-if="errormsg" class="error">{{ errormsg }}</div>
      <div class="card" @click="goToEmployees">
            <img src="/employees.png" alt="Employees" />
            <h2 class="button-text">Employees</h2>
      </div>
  
      <div class="card" @click="goToProjects">
            <img src="/projects.png" alt="projects" />
            <h2 class="button-text">Projects</h2>
      </div>
  
      <div class="card" @click="goToDepartments">
            <img src="/departments.png" alt="Departments" />
            <h2 class="button-text">Departments</h2>
      </div>
    </div>
  </template>

<style>
.container {
    display: flex;
    flex-direction: column;
    align-items: center;        
    justify-content: center;
    padding: 20px;
    width: 100%;                
    max-width: 1500px;
    height: 100%;           
    margin: 0 auto;             
}

.card {
    display: flex;
    flex-direction: row;
    align-items: center;
    background-color: #f5f5f5;
    border-radius: 12px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 1400px;
    padding: 15px;
    margin: 10px 0;
    cursor: pointer;
    transition: transform 0.2s;
    height: 200px;
}

.card:hover {
    transform: scale(1.03);
}

.card img {
    
    margin: 30px;
}

.card-content {
    flex-grow: 1;
}

.card-button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 8px 12px;
    border-radius: 6px;
    cursor: pointer;
    margin-top: 10px;
}

.card-button:hover {
    background-color: #0056b3;
}

.error {
    color: red;
    margin-bottom: 10px;
    width: 100%;
    text-align: center;
}

.button-text {
    font-size: 35px;
    font-weight: bold;
    margin-left: 10px;
}
</style>