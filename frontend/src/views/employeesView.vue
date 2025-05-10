<script>
export default {
    data: function(){
        return{
            errormsg: null,
            employees: [],
            selectedEmployee: null,
        };
    },
    methods: {
        async getEmployees() {
            this.errormsg = "";
            if (this.selectedEmployee.length === 0) {
                this.listUsers = [];
                return;
            }
            try {
                let response = await this.$axios.get(
                    `profiles/${sessionStorage.username}/employee?query=${this.selectedEmployee}`,
                    { headers: { Authorization: sessionStorage.token } }
                );
                this.listUsers = response.data || [];
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
    }
}
</script>

<template>
    <div class="employees-container">
        <h1 class="title">Employees</h1>
        <hr class="divider" />
  
        <div class="search-bar">
            <input
            v-model="searchQuery"
            type="text"
            placeholder="Search employees..."
            class="search-input"
            />
            <button class="search-button" @click="getEmployees">Search</button>
        </div>
    
        <div v-if="errormsg" class="error">{{ errormsg }}</div>
    </div>
</template>  

<style>
.employees-container {
  width: 100%;
  margin: 0 auto;
  padding: 0 20px;
  display: flex;
  flex-direction: column;        
  justify-content: flex-start; 
  height: 100%;
}


.title {
    font-size: 32px;
    font-weight: bold;
    text-align: center;
    margin-bottom: 10px;
    font-family: Arial, sans-serif;
    padding-top: 0;
}

.divider {
    border: none;
    border-top: 2px solid #ccc;
    margin-bottom: 20px;
    width: 100%;
}

.search-bar {
    width: 100%;
    display: flex;
    justify-content: center;
    gap: 10px;
    margin-bottom: 20px;
}

.search-input {
    width: 100%;
    max-width: 400px;
    padding: 10px 15px;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 16px;
}

.search-button {
    padding: 10px 20px;
    background-color: #007bff;
    border: none;
    border-radius: 8px;
    color: white;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
}

.search-button:hover {
    background-color: #0056b3;
}

.error {
    color: red;
    text-align: center;
    margin-top: 10px;
}

</style>
