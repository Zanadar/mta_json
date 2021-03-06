package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Result struct {
		XMLName   xml.Name `xml:"service"`
		Timestamp string   `xml:"timestamp"`
	}

	v := Result{}

	err := xml.Unmarshal([]byte(data()), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	output, err := json.Marshal(v)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	os.Stdout.Write(output)

	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Timestamp: %q\n", v.Timestamp)
}

func data() string {
	return `
 <service><responsecode>0</responsecode><timestamp>10/25/2016 6:38:00 AM</timestamp><subway>
  <line>
    <name>123</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>456</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>7</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>ACE</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>BDFM</name>
    <status>PLANNED WORK</status>
    <text>
                    &lt;span class="TitlePlannedWork" &gt;Planned Work&lt;/span&gt;
                    &lt;br/&gt;
                  &lt;a class="plannedWorkDetailLink" onclick=ShowHide(135741);&gt;
&lt;b&gt;[F] Some northbound trains skip 75 Av, Briarwood, Sutphin Blvd &lt;i&gt;and &lt;/i&gt;terminate at Parsons Blvd
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 135741 class="plannedWorkDetail" &gt;Rush Hours&lt;/b&gt;, 6 AM to 10 AM &lt;i&gt;and&lt;/i&gt; 3:30 PM to 8:30 PM, Mon to Fri, until Nov 18
&lt;br&gt;
&lt;br&gt;For service &lt;i&gt;to&lt;/i&gt; skipped stations, 169 St and 179 St, transfer to a Jamaica-bound [F] at 71 Av,
&lt;br&gt;Union Tpke, or Parsons Blvd.
&lt;br&gt;
&lt;br&gt;&amp;bull; Please allow additional travel time for these stations.
&lt;br&gt;&lt;b&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/25/2016</Date>
    <Time> 6:32AM</Time>
  </line>
  <line>
    <name>G</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>JZ</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>L</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>NQR</name>
    <status>SERVICE CHANGE</status>
    <text>
                    &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                    &lt;span class="DateStyle"&gt;
                    &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:21AM
                    &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;
                  &lt;P&gt;Due to NYPD activity at &lt;STRONG&gt;Newkirk Plaza&lt;/STRONG&gt;,Â northbound [Q] trains are runningÂ express from &lt;STRONG&gt;Kings Hwy&lt;/STRONG&gt; toÂ &lt;STRONG&gt;Prospect Park&lt;/STRONG&gt;.&lt;/P&gt;
&lt;P&gt;Allow additional travel time.&lt;/P&gt;
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/25/2016</Date>
    <Time> 6:21AM</Time>
  </line>
  <line>
    <name>S</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>SIR</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
</subway><bus>
  <line>
    <name>B1 - B84</name>
    <status>PLANNED WORK</status>
    <text>
                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(117571);&gt;
&lt;b&gt;B9 buses may experience delays due to track work on the N line between Stillwell Av and 8 Av
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 117571 class="plannedWorkDetail" &gt;&lt;/b&gt;Until further notice
&lt;br&gt;
&lt;br&gt;&amp;#149; Please allow additional travel time.
&lt;br&gt;&lt;b&gt;          &lt;/b&gt;&lt;/font&gt;&lt;a href=http://bustime.mta.info/#b9 target=_blank&gt;&lt;img height=30px src=http://web.mta.info/nyct/images/Bus_Time_logo.png&gt;&lt;/a&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;

                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(132028);&gt;
&lt;b&gt;B35/B35&lt;/b&gt; LTD&lt;b&gt;, B63 and B70 buses may experience delays due to construction on 5 Av at 39 St
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 132028 class="plannedWorkDetail" &gt;&lt;/b&gt;Until further notice
&lt;br&gt;
&lt;br&gt;&amp;#149; &lt;/font&gt;Please allow additional travel time.
&lt;br&gt;&lt;b&gt;          &lt;/b&gt;&lt;a href=http://bustime.mta.info/#b9 target=_blank&gt;&lt;img height=30px src=http://web.mta.info/nyct/images/Bus_Time_logo.png&gt;&lt;/a&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:32AM</Time>
  </line>
  <line>
    <name>B100 - B103</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>BM1 - BM5</name>
    <status>PLANNED WORK</status>
    <text>
                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(138552);&gt;
&lt;b&gt;QM2, QM4, QM5, QM6, QM15 and BM5 buses experienceing delays in the Queens Midtown Tunnel
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 138552 class="plannedWorkDetail" &gt;&lt;/b&gt;Until Spring 2018
&lt;br&gt;
&lt;br&gt;Due to Superstorm Sandy restoration and repair work, please allow additional travel time.
&lt;br&gt;&lt;b&gt;          &lt;/b&gt;&lt;a href=http://bustime.mta.info/#b9 target=_blank&gt;&lt;img height=30px src=http://web.mta.info/nyct/images/Bus_Time_logo.png&gt;&lt;/a&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:32AM</Time>
  </line>
  <line>
    <name>BX1 - BX55</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>BXM1 - BXM18</name>
    <status>PLANNED WORK</status>
    <text>
                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(133109);&gt;
&lt;b&gt;BxM10 stop on Morris Park Av at Eastchester Rd temporarily closed
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 133109 class="plannedWorkDetail" &gt;&lt;/b&gt;Until Oct 2018
&lt;br&gt;
&lt;br&gt;Due to construction, please use the nearby stop &lt;b&gt;on Eastchester Rd &lt;/b&gt;at Morris Park Av in both directions.
&lt;br&gt;&lt;b&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:32AM</Time>
  </line>
  <line>
    <name>M1 - M116</name>
    <status>SERVICE CHANGE</status>
    <text>
                  &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                  &lt;span class="DateStyle"&gt;
                  &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:15AM
                  &lt;/span&gt;
                  &lt;br/&gt;
                  &lt;br/&gt;
                &lt;P&gt;&lt;STRONG&gt;M79&lt;/STRONG&gt; buses are detoured in both directions, due to paving on 79 St between 3 Av and 5 Av.&lt;/P&gt;
&lt;P&gt;Detour is as follows:&lt;/P&gt;
&lt;P&gt;&lt;STRONG&gt;Eastbound:&lt;/STRONG&gt; Via 79 St, right on 5Â Av, left on 72 St,Â left on 3 Av, right on 79 St and regular route. &lt;/P&gt;
&lt;P&gt;&lt;STRONG&gt;Westbound:&lt;/STRONG&gt; ViaÂ 79 St, right on 3 Av, left on 86 St, left on 5 Av, right on 79 St and regular route. &lt;/P&gt;
&lt;P&gt;Corresponding stops will be made along the detoured route. &lt;/P&gt;
&lt;P&gt;Allow additional travel time. &lt;/P&gt;
              &lt;br/&gt;
              &lt;br/&gt;

                  &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                  &lt;span class="DateStyle"&gt;
                  &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:14AM
                  &lt;/span&gt;
                  &lt;br/&gt;
                  &lt;br/&gt;
                &lt;P&gt;&lt;STRONG&gt;M8&lt;/STRONG&gt; buses are detoured, due to Con Ed work on Christopher St between Greenwich Av and Waverly Pl. &lt;/P&gt;
&lt;P&gt;Detour is as follows:&lt;/P&gt;
&lt;P&gt;&lt;STRONG&gt;M8 Westbound&lt;/STRONG&gt;: Via 9 St, right on Greenwich Av, left on 7 Av, right on Christopher St and regular route.&lt;/P&gt;
&lt;P&gt;Corresponding stops will be made along detour route.&lt;/P&gt;
&lt;P&gt;Allow additional travel time.&lt;/P&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:15AM</Time>
  </line>
  <line>
    <name>Q1 - Q113</name>
    <status>SERVICE CHANGE</status>
    <text>
                  &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                  &lt;span class="DateStyle"&gt;
                  &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:15AM
                  &lt;/span&gt;
                  &lt;br/&gt;
                  &lt;br/&gt;
                &lt;P&gt;&lt;STRONG&gt;Q11, Q21, Q52-LTD, Q53-LTD, QM15, QM16Â &lt;/STRONG&gt;and&lt;STRONG&gt; QM17&lt;/STRONG&gt; buses are detoured due to DOT work on Woodhaven Blvd from Furmanville Av to Metropolitan Av. &lt;/P&gt;
&lt;P&gt;Detour is as follows: &lt;/P&gt;
&lt;P&gt;&lt;STRONG&gt;Southbound&lt;/STRONG&gt;: Via Woodhaven Blvd, right on Furmanville St, left on 80 St, left on Metropolitan Av, right on Woodhaven Blvd and regular route. &lt;/P&gt;
&lt;P&gt;Allow additional travel time. &lt;/P&gt;
              &lt;br/&gt;
              &lt;br/&gt;

                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(68484);&gt;
&lt;b&gt;Q69 - Southbound stop on 28 St between Queens Plaza South and 42 Rd temporarily relocated
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 68484 class="plannedWorkDetail" &gt;&lt;/b&gt;Until further notice
&lt;br&gt;
&lt;br&gt;Due to construction, please use the nearby temporary stop &lt;b&gt;on 42 Rd&lt;/b&gt; between Jackson Av and 28 St.
&lt;br&gt;&lt;/font&gt;&lt;b&gt;                                 [TP]
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;

                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(137403);&gt;
&lt;b&gt;Q25/Q25&lt;/b&gt; LTD -&lt;b&gt; No service on 5 Av at  College Point Blvd to 119 St&lt;br clear=left&gt;and 119 St at 5 Av to College Point Blvd at 7 Av
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 137403 class="plannedWorkDetail" &gt;&lt;/b&gt;6:30 AM to 5 PM, Mon to Fri, until further notice
&lt;br&gt;
&lt;br&gt;Due to construction:
&lt;br&gt;
&lt;br&gt;&lt;b&gt;Northbound&lt;/b&gt;: Buses run via 127 St.
&lt;br&gt;
&lt;br&gt;&lt;b&gt;Southbound&lt;/b&gt;: Buses stops will not be made on 5 Av, 119 St or 9 Av.  Please board buses on College Point Blvd at 7 Av.
&lt;br&gt;
&lt;br&gt;&lt;b&gt;&lt;i&gt;&lt;a style="cursor:pointer; text-decoration:underline;" onclick=ShowHide(1374030);&gt;Show Reroute Details&lt;/a&gt;&lt;/i&gt;&lt;/b&gt;
&lt;br&gt;&lt;div id="1374030"; style="display:none;"&gt;&lt;b&gt;Northbound
&lt;br&gt;&lt;/b&gt;Via 127 St
&lt;br&gt;Left on 5 Av at College Point Blvd (last/first stop)
&lt;br&gt;
&lt;br&gt;&lt;b&gt;Southbound
&lt;br&gt;&lt;/b&gt;Via 5 Av
&lt;br&gt;South on College Point Blvd
&lt;br&gt;Left on 7 Av then regular route&lt;/div&gt;
&lt;br&gt;&lt;b&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:15AM</Time>
  </line>
  <line>
    <name>QM1 - QM44</name>
    <status>SERVICE CHANGE</status>
    <text>
                  &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                  &lt;span class="DateStyle"&gt;
                  &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:15AM
                  &lt;/span&gt;
                  &lt;br/&gt;
                  &lt;br/&gt;
                &lt;P&gt;&lt;STRONG&gt;Q11, Q21, Q52-LTD, Q53-LTD, QM15, QM16Â &lt;/STRONG&gt;and&lt;STRONG&gt; QM17&lt;/STRONG&gt; buses are detoured due to DOT work on Woodhaven Blvd from Furmanville Av to Metropolitan Av. &lt;/P&gt;
&lt;P&gt;Detour is as follows: &lt;/P&gt;
&lt;P&gt;&lt;STRONG&gt;Southbound&lt;/STRONG&gt;: Via Woodhaven Blvd, right on Furmanville St, left on 80 St, left on Metropolitan Av, right on Woodhaven Blvd and regular route. &lt;/P&gt;
&lt;P&gt;Allow additional travel time. &lt;/P&gt;
              &lt;br/&gt;
              &lt;br/&gt;

                  &lt;span class="TitlePlannedWork" &gt;Planned Detour&lt;/span&gt;
                  &lt;br/&gt;
                &lt;a class="plannedWorkDetailLink" onclick=ShowHide(138552);&gt;
&lt;b&gt;QM2, QM4, QM5, QM6, QM15 and BM5 buses experienceing delays in the Queens Midtown Tunnel
&lt;/a&gt;&lt;br/&gt;&lt;br/&gt;&lt;div id= 138552 class="plannedWorkDetail" &gt;&lt;/b&gt;Until Spring 2018
&lt;br&gt;
&lt;br&gt;Due to Superstorm Sandy restoration and repair work, please allow additional travel time.
&lt;br&gt;&lt;b&gt;          &lt;/b&gt;&lt;a href=http://bustime.mta.info/#b9 target=_blank&gt;&lt;img height=30px src=http://web.mta.info/nyct/images/Bus_Time_logo.png&gt;&lt;/a&gt;
&lt;br&gt;&lt;/div&gt;&lt;/b&gt;&lt;br/&gt;
              &lt;br/&gt;
              &lt;br/&gt;
            </text>
    <Date>10/25/2016</Date>
    <Time> 6:15AM</Time>
  </line>
  <line>
    <name>S40 - S98</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>x1 - x68</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
</bus><BT>
  <line>
    <name>Bronx-Whitestone</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Cross Bay</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Henry Hudson</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Hugh L. Carey</name>
    <status>SERVICE CHANGE</status>
    <text>
                    &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                    &lt;span class="DateStyle"&gt;
                    &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:12AM
                    &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;
                  HLC - HOV Lane Open 6 AM to 10 AM. Two-Way Operations in effect. Three (3) lanes Manhattan bound. One (1) lane Brooklyn bound.
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/25/2016</Date>
    <Time> 6:12AM</Time>
  </line>
  <line>
    <name>Marine Parkway</name>
    <status>SERVICE CHANGE</status>
    <text>
                    &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                    &lt;span class="DateStyle"&gt;
                    &amp;nbsp;Posted:&amp;nbsp;10/24/2016&amp;nbsp; 6:23AM
                    &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;
                  Marine Parkway-Gil Hodges Memorial Bridge; One lane closed Rockaway-bound until December 2016. Motorists should allow extra time or use alternate route.
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/24/2016</Date>
    <Time> 6:23AM</Time>
  </line>
  <line>
    <name>Queens Midtown</name>
    <status>SERVICE CHANGE</status>
    <text>
                    &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                    &lt;span class="DateStyle"&gt;
                    &amp;nbsp;Posted:&amp;nbsp;10/25/2016&amp;nbsp; 6:11AM
                    &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;
                  QMT - HOV Lane Open 6 AM to 10 AM. Two-Way Operation in effect. Three (3) lanes Manhattan bound. One (1) lane Queens bound.
                &lt;br/&gt;&lt;br/&gt;

                    &lt;span class="TitlePlannedWork" &gt;Planned Work&lt;/span&gt;
                    &lt;br/&gt;
                  Queens-Midtown Tunnel downtown exit; One lane closed. Use 37th St tunnel exit for access to 2nd Ave. Motorists should allow extra time and may wish to use an alternate route if possible" Drivers should expect delays and plan accordingly. Motorists can sign up for MTA e-mail or text alerts at &lt;A href="http://www.mta.info"&gt;www.mta.info&lt;/A&gt; and check the Bridges and Tunnels homepage or Facebook page for the latest information on this planned work.
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/25/2016</Date>
    <Time> 6:11AM</Time>
  </line>
  <line>
    <name>Robert F. Kennedy</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Throgs Neck</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Verrazano-Narrows</name>
    <status>PLANNED WORK</status>
    <text>
                    &lt;span class="TitlePlannedWork" &gt;Planned Work&lt;/span&gt;
                    &lt;br/&gt;
                  VNB: PLANNED WORK; S. I. BOUND LOWER LEVEL - ONE LANE CLOSED; EXPECT DELAYS.
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/25/2016</Date>
    <Time> 5:51AM</Time>
  </line>
</BT><LIRR>
  <line>
    <name>Babylon</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>City Terminal Zone</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Far Rockaway</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Hempstead</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Long Beach</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Montauk</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Oyster Bay</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Port Jefferson</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Port Washington</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Ronkonkoma</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>West Hempstead</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
</LIRR><MetroNorth>
  <line>
    <name>Hudson</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Harlem</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Wassaic</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>New Haven</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>New Canaan</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Danbury</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Waterbury</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
  <line>
    <name>Pascack Valley</name>
    <status>SERVICE CHANGE</status>
    <text>
                    &lt;span class="TitleServiceChange" &gt;Service Change&lt;/span&gt;
                    &lt;span class="DateStyle"&gt;
                    &amp;nbsp;Posted:&amp;nbsp;10/22/2016&amp;nbsp; 6:14AM
                    &lt;/span&gt;&lt;br/&gt;&lt;br/&gt;
                  &lt;P&gt;Buses will replace train service on the Pascack Valley Line to accommodate grade crossing repairs and maintenance work weekends from October 8 to November 5 and midday weekdays from October 10 to November 4. &lt;/P&gt;
&lt;P&gt;Weekends of October 22-23 and 29-30 No trains operate. Local and New York State buses operate between Spring Valley and Secaucus. &lt;/P&gt;
&lt;P&gt;Weekday midday hours from Monday, October 10 - November 4 Buses will operate between Spring Valley and New Bridge Landing during midday hours only. &lt;/P&gt;
&lt;P&gt;See full details at http://bit.ly/2e6uYpE&lt;/P&gt;
                &lt;br/&gt;&lt;br/&gt;
              </text>
    <Date>10/22/2016</Date>
    <Time> 6:14AM</Time>
  </line>
  <line>
    <name>Port Jervis</name>
    <status>GOOD SERVICE</status>
    <text />
    <Date></Date>
    <Time></Time>
  </line>
</MetroNorth></service>


 `
}
