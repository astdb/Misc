
/*
You are given a data structure of employee information, which includes the employee's unique id, his importance value and his direct subordinates' id.

For example, employee 1 is the leader of employee 2, and employee 2 is the leader of employee 3. They have importance value 15, 10 and 5, respectively. Then employee 1 has a data structure like [1, 15, [2]], and employee 2 has [2, 10, [3]], and employee 3 has [3, 5, []]. Note that although employee 3 is also a subordinate of employee 1, the relationship is not direct.

Now given the employee information of a company, and an employee id, you need to return the total importance value of this employee and all his subordinates.

Example 1:
Input: [[1, 5, [2, 3]], [2, 3, []], [3, 3, []]], 1
Output: 11
Explanation:
Employee 1 has importance value 5, and he has two direct subordinates: employee 2 and employee 3. They both have importance value 3. So the total importance value of employee 1 is 5 + 3 + 3 = 11.
*/

/*
// Employee info
class Employee {
    // It's the unique id of each node;
    // unique id of this employee
    public int id;
    // the importance value of this employee
    public int importance;
    // the id of direct subordinates
    public List<Integer> subordinates;
};
*/

import java.util.*;

public class EmployeeImportance {
    public static void main(String[] args) {
        List<Employee> empList = new ArrayList<Employee>();
        
        List<Integer> emp1Subords = new ArrayList<Integer>();
        emp1Subords.add(2);
        emp1Subords.add(3);
        Employee emp1 = new Employee(1,5,emp1Subords);
        empList.add(emp1);

        List<Integer> emp2Subords = new ArrayList<Integer>();
        Employee emp2 = new Employee(2,3,emp2Subords);
        empList.add(emp2);

        List<Integer> emp3Subords = new ArrayList<Integer>();
        Employee emp3 = new Employee(3,3,emp3Subords);
        empList.add(emp3);

        System.out.println(getImportance(empList, 1));
    }

    public static int getImportance(List<Employee> employees, int id) {
        // transform employee list into hashmap
        Map<Integer, Employee> empMap = new HashMap<Integer, Employee>();
        for(Employee emp: employees) {
            empMap.put(emp.id, emp);
        }

        return importanceTotal(empMap, id);
    }

    public static int importanceTotal(Map<Integer,Employee> employees, int id) {
        Employee thisEmp = employees.get(id);

        if(thisEmp.subordinates.size() == 0) {
            return thisEmp.importance;

        } else {
            int subsImpTot = 0;
            for(Integer subID: thisEmp.subordinates) {
                subsImpTot += importanceTotal(employees, subID);
            }

            return thisEmp.importance + subsImpTot;
        }
    }
}

class Employee {
    Employee(int id, int importance, List<Integer> subordinates) {
        this.id = id;
        this.importance = importance;
        this.subordinates = subordinates;
    }
    // It's the unique id of each node;
    // unique id of this employee
    public int id;
    // the importance value of this employee
    public int importance;
    // the id of direct subordinates
    public List<Integer> subordinates;
};
