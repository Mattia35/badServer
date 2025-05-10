<script>
export default {
    data: function(){
        return{
            errormsg: null,
            employees: [],
            selectedEmployee: "",
        };
    },
    methods: {
        async getEmployees() {
            this.errormsg = "";
            if (this.selectedEmployee.length === 0) {
                this.employees = [];
                return;
            }
            try {
                let response = await this.$axios.get(
                    `profiles/${sessionStorage.username}/employees?query=${this.selectedEmployee}`,
                    { headers: { Session: sessionStorage.session, Token: sessionStorage.token } }
                );
                this.employees = response.data || [];
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
            v-model="selectedEmployee"
            type="text"
            placeholder="Search employees..."
            class="search-input"
            />
            <button class="search-button" @click="getEmployees">Search</button>
        </div>
    
        <div v-if="errormsg" class="error">{{ errormsg }}</div>

        <div class="employees-list">
            <div
                v-for="employee in employees"
                class="employee-box"
            >
                <h2 class="employee-title">{{ employee.name_surname }}</h2>
                <p><strong>Email:</strong> {{ employee.email }}</p>
                <p><strong>Phone:</strong> {{ employee.phone }}</p>
                <p><strong>Department:</strong> {{ employee.department }}</p>
                <p><strong>Position:</strong> {{ employee.position }}</p>
                <p><strong>Project ID:</strong> {{ employee.project }}</p>
            </div>
        </div>
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
    max-width: 800px;
    margin-left: auto;
    margin-right: auto;
}

.search-input {
    width: 100%;
    padding: 10px 15px;
    border: 1px solid #ccc;
    border-radius: 8px;
    font-size: 16px;
    outline: none;
    box-shadow: none;
}

.search-input:focus {
    outline: none;
    box-shadow: none;
    border-color: #ccc; 
}

.search-button {
    padding: 10px 15px;
    background-color: #12B886;
    border: none;
    border-radius: 8px;
    color: white;
    font-size: 16px;
    cursor: pointer;
    transition: background-color 0.2s ease;
}

.employee-list {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.employee-box {
    border: 1px solid #ccc;
    border-radius: 12px;
    padding: 20px;
    background-color: #f9f9f9;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.05);
}

.employee-title {
    font-size: 20px;
    font-weight: bold;
    margin-bottom: 10px;
}

.error {
    color: red;
    text-align: center;
    margin-top: 10px;
}

</style>
