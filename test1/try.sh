export GOVC_INSECURE=1
export GOVC_URL='https://root:LambShank5612@10.19.114.57/sdk'
echo $GOVC_URL
govc ls /ha-datacenter/vm
govc vm.info vm/ctl.k.e2e.bos.redhat.com/
govc vm.info 564d6f78-2174-0a4a-240e-5d23173a0ea0
govc device.ls -vm vm/ctl.k.e2e.bos.redhat.com
#govc  object.collect -json vm/ctl.k.e2e.bos.redhat.com guest.net | jq
