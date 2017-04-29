package main

import (
	"encoding/json"
	"fmt"
	"time"

	swl "github.com/stohio/software-lab/lib"
)

var currentNodeId int
var currentNetworkId int
var currentStackId int

var nodes swl.Nodes
var networks swl.Networks

var stacks swl.Stacks

func init() {

	var stack swl.Stack

	jsonStack := `{
  "id": 1,
  "name": "Hackathon",
  "softwares": [
    {
      "id": 1,
      "name": "Node JS",
      "publisher": "Node",
      "versions": [
        {
          "id": 1,
          "version": "6.10.0",
          "os": "Windows",
          "architecture": "32",
          "extension": ".msi",
          "url": "https://nodejs.org/dist/v6.10.0/node-v6.10.0-x86.msi"
        },
        {
          "id": 2,
          "version": "6.10.0",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "https://nodejs.org/dist/v6.10.0/node-v6.10.0.pkg"
        }
      ]
    },
    {
      "id": 2,
      "name": "Unity",
      "publisher": "Unity",
      "versions": [
        {
          "id": 1,
          "version": "5.5.2",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "http://netstorage.unity3d.com/unity/3829d7f588f3/MacEditorInstaller/Unity-5.5.2f1.pkg"
        },
        {
          "id": 2,
          "version": "5.5.2",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "http://netstorage.unity3d.com/unity/3829d7f588f3/Windows32EditorInstaller/UnitySetup32-5.5.2f1.exe"
        }
      ]
    },
    {
      "id": 3,
      "name": "Arduino IDE",
      "publisher": "Arduino",
      "versions": [
        {
          "id": 1,
          "version": "1.8.1",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://downloads.arduino.cc/arduino-1.8.1-windows.exe"
        },
        {
          "id": 2,
          "version": "1.8.1",
          "os": "Mac",
          "architecture": "64",
          "extension": ".zip",
          "url": "https://downloads.arduino.cc/arduino-1.8.1-macosx.zip"
        }
      ]
    },
    {
      "id": 4,
      "name": "Code Blocks",
      "publisher": "Code Blocks",
      "versions": [
        {
          "id": 1,
          "version": "13.12",
          "os": "Mac",
          "architecture": "64",
          "extension": ".zip",
          "url": "https://superb-dca2.dl.sourceforge.net/project/codeblocks/Binaries/13.12/MacOS/CodeBlocks-13.12-mac.zip"
        },
        {
          "id": 2,
          "version": "16.01",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://svwh.dl.sourceforge.net/project/codeblocks/Binaries/16.01/Windows/codeblocks-16.01mingw-setup.exe"
        }
      ]
    },
    {
      "id": 5,
      "name": "Brackets",
      "publisher": "Brackets",
      "versions": [
        {
          "id": 1,
          "version": "1.9",
          "os": "Mac",
          "architecture": "64",
          "extension": ".dmg",
          "url": "https://github-cloud.s3.amazonaws.com/releases/2935735/f72fa2b2-04d7-11e7-8cc6-0ff6d0b5c052.dmg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAISTNZFOVBIJMK3TQ%2F20170311%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170311T145412Z&X-Amz-Expires=300&X-Amz-Signature=8d53cf32d796b72badac135ebb09f342d899028a75e313033cf4c83cd19e9e64&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3DBrackets.Release.release-1.9-prerelease.dmg&response-content-type=application%2Foctet-stream"
        },
        {
          "id": 2,
          "version": "1.9",
          "os": "Windows",
          "architecture": "32",
          "extension": ".msi",
          "url": "https://github-cloud.s3.amazonaws.com/releases/2935735/f73014c2-04d7-11e7-8e6b-358e1891a360.msi?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAISTNZFOVBIJMK3TQ%2F20170311%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20170311T145341Z&X-Amz-Expires=300&X-Amz-Signature=b2a7a2a7d0af41dd3a83834396369ff94bcd2e8b899411b64e6fcb77e6db5efa&X-Amz-SignedHeaders=host&actor_id=0&response-content-disposition=attachment%3B%20filename%3DBrackets.Release.release-1.9-prerelease.msi&response-content-type=application%2Foctet-stream"
        }
      ]
    },
    {
      "id": 6,
      "name": "Python",
      "publisher": "Python",
      "versions": [
        {
          "id": 1,
          "version": "3.6.0",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "https://www.python.org/ftp/python/3.6.1/python-3.6.1rc1-macosx10.6.pkg"
        },
        {
          "id": 2,
          "version": "3.6.0",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://www.python.org/ftp/python/3.6.0/python-3.6.0.exe"
        }
      ]
    },
    {
      "id": 7,
      "name": "NOOBS",
      "publisher": "NOOBS",
      "versions": [
        {
          "id": 1,
          "version": "2.3.0",
          "os": "Raspberry Pi",
          "architecture": "ARM",
          "extension": ".zip",
          "url": "http://director.downloads.raspberrypi.org/NOOBS/images/NOOBS-2017-03-03/NOOBS_v2_3_0.zip"
        }
      ]
    },
    {
      "id": 8,
      "name": "MAMP (Mac Apache, MySQL, PHP)",
      "publisher": "MAMP",
      "versions": [
        {
          "id": 1,
          "version": "4.1.1",
          "os": "Mac",
          "architecture": "64",
          "extension": ".pkg",
          "url": "http://downloads1.mamp.info/MAMP-PRO/releases/4.1.1/MAMP_MAMP_PRO_4.1.1.pkg"
        }
      ]
    },
    {
      "id": 9,
      "name": "WAMP (Windows Apache, MySQL, PHP)",
      "publisher": "WAMP",
      "versions": [
        {
          "id": 1,
          "version": "3.0.6",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://svwh.dl.sourceforge.net/project/wampserver/WampServer%203/WampServer%203.0.0/wampserver3.0.6_x86_apache2.4.23_mysql5.7.14_php5.6.25-7.0.10.exe"
        }
      ]
    },
    {
      "id": 10,
      "name": "Github Desktop",
      "publisher": "Github",
      "versions": [
        {
          "id": 1,
          "version": "3.3.4",
          "os": "Mac",
          "architecture": "64",
          "extension": ".zip",
          "url": "https://mac-installer.github.com/mac/GitHub%20Desktop%20222.zip"
        },
        {
          "id": 2,
          "version": "3.3.4",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://github-windows.s3.amazonaws.com/GitHubSetup.exe"
        }
      ]
    },
    {
      "id": 11,
      "name": "VS Code",
      "publisher": "Microsoft",
      "versions": [
        {
          "id": 1,
          "version": "1.10.2",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://az764295.vo.msecnd.net/stable/8076a19fdcab7e1fc1707952d652f0bb6c6db331/VSCodeSetup-1.10.2.exe"
        },
        {
          "id": 2,
          "version": "1.10.2",
          "os": "Mac",
          "architecture": "64",
          "extension": ".zip",
          "url": "https://az764295.vo.msecnd.net/stable/8076a19fdcab7e1fc1707952d652f0bb6c6db331/VSCode-darwin-stable.zip"
        }
      ]
    },
    {
      "id": 12,
      "name": "Android Studio",
      "publisher": "Google",
      "versions": [
        {
          "id": 1,
          "version": "2.3.0.8",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://dl.google.com/dl/android/studio/install/2.3.0.8/android-studio-bundle-162.3764568-windows.exe"
        },
        {
          "id": 2,
          "version": "2.3.0.8",
          "os": "Mac",
          "architecture": "64",
          "extension": ".dmg",
          "url": "https://dl.google.com/dl/android/studio/install/2.3.0.8/android-studio-ide-162.3764568-mac.dmg"
        }
      ]
    },
    {
      "id": 13,
      "name": "PyCharm Community",
      "publisher": "IntelliJ",
      "versions": [
        {
          "id": 1,
          "version": "2017",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://download-cf.jetbrains.com/python/pycharm-community-2017.1.exe"
        },
        {
          "id": 2,
          "version": "2017",
          "os": "Mac",
          "architecture": "64",
          "extension": ".dmg",
          "url": "https://download-cf.jetbrains.com/python/pycharm-community-2017.1.dmg"
        }
      ]
    },
    {
      "id": 14,
      "name": "IntelliJ Idea Community",
      "publisher": "IntelliJ",
      "versions": [
        {
          "id": 1,
          "version": "2017",
          "os": "Windows",
          "architecture": "32",
          "extension": ".exe",
          "url": "https://download-cf.jetbrains.com/idea/ideaIC-2017.1.exe"
        },
        {
          "id": 2,
          "version": "2017",
          "os": "Mac",
          "architecture": "64",
          "extension": ".dmg",
          "url": "https://download-cf.jetbrains.com/idea/ideaIC-2017.1.dmg"
        }
      ]
    }
  ]
}`

	if err := json.Unmarshal([]byte(jsonStack), &stack); err != nil {
		panic(err)
	}

	var softs swl.Softwares
	softs = append(softs, stack.Softwares[0])
	softs = append(softs, stack.Softwares[1])
	pack := swl.Package{
		Id:          1,
		Name:        "My First Package",
		Description: "This is the first package",
		Softwares:   softs,
	}
	stack.Packages = append(stack.Packages, &pack)
	stacks = append(stacks, &stack)

}

// RepoFindStack returns the stack associated with an id
// @param id: the id of the stack to get
// @return: returns a pointer to the stack with the id id
func RepoFindStack(id int) *swl.Stack {
	for _, s := range stacks {
		if s.Id == id {
			return s
		}
	}
	return nil
}

// RepoCreateStack takes a stack structure and adds this stack to the list of stacks
// @param s: a pointer to the newly created stack
// @return: returns the stack with the its ID set
func RepoCreateStack(s *swl.Stack) *swl.Stack {
	currentStackId += 1
	s.Id = currentStackId
	stacks = append(stacks, s)
	return s
}

// RepoDestroyStack removes a stack with the specified id
// @param id: the id of the stack to destroy
// @return: returns nil if successful and returns and error if it can't find the stack
func RepoDestroyStack(id int) error {
	for i, s := range stacks {
		if s.Id == id {
			stacks = append(stacks[:i], stacks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Stack with id of %d to delete", id)
}

// RepoFindNetworkByIP gets the network with the specified ip
// @param ip: the ip of the network to get
// @return: the network with the ip specified, nil if the network can't be found
func RepoFindNetworkByIP(ip string) *swl.Network {
	for _, net := range networks {
		if net.IP == ip {
			return net
		}
	}
	return nil
}

// RepoFindBestNodeInNEtworkByIP gets the node in the network specified by ip with the smallest number of clients
// @param ip: the ip of the network to search in
// @return: the node with the least amount of clients in the specified network
func RepoFindBestNodeInNetworkByIP(ip string) *swl.Node {
	net := RepoFindNetworkByIP(ip)
	if net == nil {
		fmt.Println("Could Not Find Network")
		return nil
	}
	var bestNode *swl.Node
	bestDownloads := -1
	for _, n := range net.Nodes {
		fmt.Printf("Node, Best: %d. %d\n", n.Clients, bestDownloads)
		if (n.Clients < bestDownloads || bestDownloads == -1) && (n.Enabled) {
			fmt.Println("Best Node Updated!")
			bestNode = n
			bestDownloads = n.Clients
			if bestDownloads == 0 {
				return bestNode
			}
		}
	}
	if bestDownloads == -1 {
		return nil
	}
	return bestNode

}

// RepoCreateNetwork takes in a network struct and adds it to the list of all networks
// @param n: a Network struct
// @return: the network that was just created
func RepoCreateNetwork(n *swl.Network) *swl.Network {
	currentNetworkId += 1
	n.Id = currentNetworkId
	fmt.Println("Added Network")
	networks = append(networks, n)
	return n
}

// RepoDestroyNetwork deletes a network witht eh given id
// @param id: the id of the network to destroy
// @return: nil if successful and an error if the network with the given id doesn't exist
func RepoDestroyNetwork(id int) error {
	for i, n := range networks {
		if n.Id == id {
			networks = append(networks[:i], networks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Unable to find Network with id of %d to delete", id)
}

// RepoFindNode returns a node with the given id
// @param id: the id of node to get
// @return: the node with the given ip or nil
func RepoFindNode(id int) *swl.Node {
	for _, n := range nodes {
		if n.Id == id {
			return n
		}
	}
	//Otherwise, Return Empty
	return nil
}

// RepoCreateNode adds the given node to list of nodes and sets its id value
// @param n: the new node to add
// @return: the new node with updated id
func RepoCreateNode(n *swl.Node) *swl.Node {
	currentNodeId += 1
	n.Id = currentNodeId
	n.Added = time.Now()
	nodes = append(nodes, n)
	return n
}

// RepoEnableNode sets the node with given id to enabled
// @param id: the id of the node to update
// @return: returns nil if the node isn't found, otherwise returns the updated node
func RepoEnableNode(id int) *swl.Node {
	node := RepoFindNode(id)
	if node == nil {
		return nil
	}
	node.Enabled = true
	return node
}

// DeleteNode removes the node with the given id
// @param id: the id of node to delete
// @return: returns nil if successful, otherwise returns an Error
func DeleteNode(id int) error {
	for _, n := range networks {
		for j, nod := range n.Nodes {
			if nod.Id == id {
				n.Nodes = append(n.Nodes[:j], n.Nodes[j+1:]...)
				fmt.Println(len(n.Nodes))
				if len(n.Nodes) == 0 {
					RepoDestroyNetwork(n.Id)
				}
				break
			}
		}
	}
	for i, n := range nodes {
		if n.Id == id {
			nodes = append(nodes[:i], nodes[i+1:]...)
			return nil
		}

	}
	return fmt.Errorf("Unable to find Node with id of %d to delete", id)
}

// RepoUpdateNodeClients either increments or decrements the Client field of the node with the given id
// @param id: the id of the node to update
// @param increment: if true the client field is incremented by 1, if false its decremented by 1
// @return: nil on success, otherwise an Error
func RepoUpdateNodeClients(id int, increment bool) error {
	for _, n := range nodes {
		if n.Id == id {
			if increment {
				n.Clients += 1
			} else {
				n.Clients -= 1
			}
			return nil
		}
	}
	return fmt.Errorf("Unable to find Node with id of %d to Update Clients", id)
}
