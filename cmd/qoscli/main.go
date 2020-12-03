package main

func main() {

	//TODO: Need to implement the cli

	//netqospolicy.DEBUG = true
	//
	//fmt.Println("existing policies")
	//policies, err := netqospolicy.GetAll()
	//if err != nil {
	//	fmt.Printf("error getting policies: %v", err)
	//}
	//fmt.Println(policies)
	//
	//fmt.Println("[+] CREATE POLICY FROM SCRATCH TEST ==================")
	//fmt.Println("create new policy")
	//dscp := int8(30)
	//port := uint16(80)
	//p := netqospolicy.NetQoSPolicy{
	//	Name:              "URL Match",
	//	DSCPAction:        &dscp,
	//	IPDstPortMatchCondition: &port,
	//}
	//err = p.Create()
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println("update policy to use uri match")
	//p.IPDstPortMatchCondition = nil
	//p.IPDstPortStartMatchCondition = nil
	//p.IPDstPortStartMatchCondition = nil
	//p.UriMatchCondition = "http://www.spam.com/send.aspx"
	//p.UriRecursiveMatchCondition = true
	//p.Update()
	//
	//fmt.Println("remove policy")
	//p.Remove()
	//
	//fmt.Println("[+] RETRIEVE AND UPDATE POLICY TEST ==================")
	//fmt.Println("create ftp policy")
	//p = netqospolicy.NetQoSPolicy{
	//	Name:                      "FTP",
	//	DSCPAction:                &dscp,
	//	AppPathNameMatchCondition: "ftp.exe",
	//}
	//err = p.Create()
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println("retrieve and update policy")
	//p2, err := netqospolicy.Get("FTP")
	//if err != nil {
	//	fmt.Printf("failed to get policy: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("making updates")
	//p2.AppPathNameMatchCondition = ""
	//p2.UriMatchCondition = "https://www.k3rn3l.io"
	//applied, _ := p2.IsApplied()
	//fmt.Printf("IsApplied: %v\n", applied)
	//fmt.Println("updating...")
	//p2.Update()
	//applied, _ = p2.IsApplied()
	//fmt.Printf("IsApplied: %v\n", applied)
	//p2.Remove()


}
