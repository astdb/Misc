/*
Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address, mail sent there will be forwarded to the same address without dots in the local name.  For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.  (Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored. This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.  (Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.  How many different addresses actually receive mails? 

 

Example 1:

Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails
 

Note:

1 <= emails[i].length <= 100
1 <= emails.length <= 100
Each emails[i] contains exactly one '@' character.
*/

import java.util.*;

class UniqueEmailAdds {
    public static void main(String[] args) {
        String[][] tests = new String[][]{{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}};

        for (int i = 0; i < tests.length; i++) {
            System.out.println(numUniqueEmails(tests[i]));
        }
    }

    public static int numUniqueEmails(String[] emails) {
        Map<String, Integer> uniqueLocalNames = new HashMap<String, Integer>();

        for (int i = 0; i < emails.length; i++) {
            String email = emails[i];
            if(email.indexOf('@') != -1) {
                String[] email_components = email.split("@");

                if(email_components.length == 2) {
                    String localname = email_components[0];
                    String domain = email_components[1];

                    // if there's a plus sign in local name, discard part after first '+'
                    if(localname.indexOf('+') != -1) {
                        String[] localname_comps = localname.split("\\+");
                        localname = localname_comps[0];
                    }

                    // if there are dots in local name, remove them
                    localname = localname.replace(".", "");
                    email = localname + "+" + domain;

                    if(!uniqueLocalNames.containsKey(email)) {
                        uniqueLocalNames.put(email, 1);
                    }

                } else {
                    // not a valid email address
                }
            } else {
                // not a valid email address

            }
        }

        return uniqueLocalNames.size();
    }
}
