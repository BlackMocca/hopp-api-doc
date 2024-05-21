

- [**AD Accounts Backend API**](/?id=folder-ad-accounts-backend-api)

  
  

  
  
  

  - [[GET] Get User](/?id=get-get-user)

  

  - [[GET] Get List User](/?id=get-get-list-user)

  

  - [[POST] Create User](/?id=post-create-user)

  

  - [[POST] Update User Info](/?id=post-update-user-info)

  

  - [[POST] Reset Password](/?id=post-reset-password)

  

  - [[POST] Change Password](/?id=post-change-password)

  

  - [[GET] List Locked Out Users](/?id=get-list-locked-out-users)

  

  - [[POST] Unlock User](/?id=post-unlock-user)

  

  - [[POST] Lock User](/?id=post-lock-user)

  

  - [[POST] Disable User](/?id=post-disable-user)

  

  - [[POST] Enable User](/?id=post-enable-user)

  

  - [[DELETE] Delete User](/?id=delete-delete-user)

  

  - [[POST] 2Factor Enable](/?id=post-2factor-enable)

  

  - [[POST] 2Factor Disable](/?id=post-2factor-disable)

  

  - [[GET] 2Factor Get Secret (JSON)](/?id=get-2factor-get-secret-json)

  

  - [[GET] 2Factor Get Secret (QRCode)](/?id=get-2factor-get-secret-qrcode)

  

  - [[POST] 2Factor Test Check](/?id=post-2factor-test-check)

  

  - [[GET] New Request](/?id=get-new-request)

  

  - [[GET] New Request](/?id=get-new-request)

  
  



- [**API**](/?id=folder-api)

  
  
  

  - **AUTH**
    
    

    - [[POST] /v1/signin](/?id=post-v1signin)

    

    - [[GET] /v1/callback](/?id=get-v1callback)

    

    - [[POST] /v1/sign/token](/?id=post-v1signtoken)

    

    - [[POST] /v1/signout](/?id=post-v1signout)

    

    - [[POST] /v1/resign/token](/?id=post-v1resigntoken)

    

    - [[GET] /v1/roles/config](/?id=get-v1rolesconfig)

    

    - [[POST] /v1/create/user/role](/?id=post-v1createuserrole)

    

    - [[PUT] /v1/edit/user/:user_id/role](/?id=put-v1edituseruser_idrole)

    

    - [[PUT] /v1/status/user/:user_id/role](/?id=put-v1statususeruser_idrole)

    

    - [[DELETE] /v1/delete/user/:user_id/role](/?id=delete-v1deleteuseruser_idrole)

    
    

  

  - **ORGANIZATION**
    
    

    - [[GET] /v1/organizes](/?id=get-v1organizes)

    

    - [[GET] /v1/organize/:organize_id](/?id=get-v1organizeorganize_id)

    

    - [[GET] /v1/organize/:organize_id/departments](/?id=get-v1organizeorganize_iddepartments)

    

    - [[GET] /v1/buildings](/?id=get-v1buildings)

    
    

  

  - **USER**
    

    - **consent**
      

        - [[POST] /v1/management/consent](/?id=post-v1managementconsent)

      

        - [[GET] /v1/management/consent](/?id=get-v1managementconsent)

      

        - [[GET] /v1/management/:consent_id/consent](/?id=get-v1managementconsent_idconsent)

      

        - [[PUT] /v1/management/status/:consent_id/consent](/?id=put-v1managementstatusconsent_idconsent)

      

        - [[GET] /v1/consent](/?id=get-v1consent)

      

        - [[PUT] /v1/consent](/?id=put-v1consent)

      

        - [[GET] /v1/my/consent](/?id=get-v1myconsent)

      
    

    - [[GET] /v1/me](/?id=get-v1me)

    

    - [[GET] /v1/types](/?id=get-v1types)

    

    - [[GET] /v1/lists](/?id=get-v1lists)

    

    - [[GET] /v1/:user_id/details](/?id=get-v1user_iddetails)

    

    - [[POST] /v1/accounts](/?id=post-v1accounts)

    

    - [[PUT] /v1/:user_id/accounts](/?id=put-v1user_idaccounts)

    

    - [[POST] /v1/:user_id/accounts](/?id=post-v1user_idaccounts)

    

    - [[POST] /v1/account/reset/passwords](/?id=post-v1accountresetpasswords)

    

    - [[POST] /v1/:user_id/account/request/reset/passwords](/?id=post-v1user_idaccountrequestresetpasswords)

    

    - [[DELETE] /v1/:user_id](/?id=delete-v1user_id)

    

    - [[POST] /v1/account/management](/?id=post-v1accountmanagement)

    

    - [[POST] /v1/account/request/reset/passwords](/?id=post-v1accountrequestresetpasswords)

    

    - [[POST] /v1/reset/password/check/ref_code](/?id=post-v1resetpasswordcheckref_code)

    

    - [[POST] /v1/reset/password](/?id=post-v1resetpassword)

    

    - [[GET] /v1/list/detail](/?id=get-v1listdetail)

    

    - [[POST] /v1/users/file](/?id=post-v1usersfile)

    

    - [[GET] /v1/users/role/list](/?id=get-v1usersrolelist)

    

    - [[GET] New Request](/?id=get-new-request)

    
    

  

  - **FILE**
    
    

    - [[POST] /v1/uploads](/?id=post-v1uploads)

    

    - [[POST] /v1/attach/uploads](/?id=post-v1attachuploads)

    

    - [[POST] /v1/attach/delete](/?id=post-v1attachdelete)

    
    

  

  - **FACILITY**
    
    

    - [[GET] /v1/global/configs](/?id=get-v1globalconfigs)

    

    - [[GET] /v1/parking/penalities](/?id=get-v1parkingpenalities)

    

    - [[GET] /v1/parking/permit/types](/?id=get-v1parkingpermittypes)

    

    - [[GET] /v1/parking/permit/:parking_id](/?id=get-v1parkingpermitparking_id)

    

    - [[GET] /v1/app/parking/permit/:parking_id](/?id=get-v1appparkingpermitparking_id)

    

    - [[GET] /v1/my/services ที่รวมใบอนุญาตของตัวเอง](/?id=get-v1myservices-thiirwmaib-nuyaatkh-ngtawe-ng)

    

    - [[GET] /v1/parking/permits](/?id=get-v1parkingpermits)

    

    - [[GET] /v1/parking/province](/?id=get-v1parkingprovince)

    

    - [[GET] /v1/parking/status](/?id=get-v1parkingstatus)

    

    - [[GET] /v1/parking/fueltype](/?id=get-v1parkingfueltype)

    

    - [[GET] v1/parking/permit/:parking_id/car](/?id=get-v1parkingpermitparking_idcar)

    

    - [[GET] v1/app/parking/permit/:parking_id/car](/?id=get-v1appparkingpermitparking_idcar)

    

    - [[POST] /v1/parking/permits](/?id=post-v1parkingpermits)

    

    - [[POST] /v1/parking/permits/:parking_id/car](/?id=post-v1parkingpermitsparking_idcar)

    

    - [[POST] /v1/app/parking/permits/:parking_id/car](/?id=post-v1appparkingpermitsparking_idcar)

    

    - [[PUT] /v1/parking/permit/:parking_id/status](/?id=put-v1parkingpermitparking_idstatus)

    

    - [[GET] /v1/gates](/?id=get-v1gates)

    

    - [[GET] /v1/parking/permit/status](/?id=get-v1parkingpermitstatus)

    

    - [[POST] /v1/parking/permits/list](/?id=post-v1parkingpermitslist)

    

    - [[GET] /v1/receipt/:parking_id](/?id=get-v1receiptparking_id)

    

    - [[GET] /v1/receipt/:parking_id/file](/?id=get-v1receiptparking_idfile)

    

    - [[GET] /v1/parking/permits/list/excel](/?id=get-v1parkingpermitslistexcel)

    

    - [[GET] /v1/parking/permits/stat/excel](/?id=get-v1parkingpermitsstatexcel)

    

    - [[GET] /v1/parking/permits - Duplicate](/?id=get-v1parkingpermits-duplicate)

    
    

  

  - **Internal**
    

    - **SCHEDULER**
      

        - [[POST] sync all organize kmutt](/?id=post-sync-all-organize-kmutt)

      

        - [[POST] publish notification timer](/?id=post-publish-notification-timer)

      

        - [[GET] publish notification timer Copy](/?id=get-publish-notification-timer-copy)

      
    

  

  - **NOTIFICATION**
    

    - **announcements**
      

        - [[GET] /v1/management/announcements](/?id=get-v1managementannouncements)

      

        - [[GET] /v1/management/:announcement_id/announcements](/?id=get-v1managementannouncement_idannouncements)

      

        - [[POST] /v1/management/announcements](/?id=post-v1managementannouncements)

      

        - [[PUT] /v1/management/:announcement_id/announcements](/?id=put-v1managementannouncement_idannouncements)

      

        - [[DELETE] /v1/management/:announcement_id/announcements](/?id=delete-v1managementannouncement_idannouncements)

      

        - [[GET] /v1/announcements](/?id=get-v1announcements)

      

        - [[GET] /v1/:announcement_id/announcements](/?id=get-v1announcement_idannouncements)

      
    

    - [[GET] /v1/notification](/?id=get-v1notification)

    

    - [[GET] /v1/notification/details](/?id=get-v1notificationdetails)

    

    - [[GET] /v1/notification/view](/?id=get-v1notificationview)

    

    - [[GET] /v1/my/notification](/?id=get-v1mynotification)

    

    - [[POST] /v1/notification](/?id=post-v1notification)

    

    - [[DELETE] /v1/my/notification](/?id=delete-v1mynotification)

    

    - [[DELETE] /v1/notification](/?id=delete-v1notification)

    

    - [[PUT] /v1/notification/:notification_id](/?id=put-v1notificationnotification_id)

    
    

  
  

  
  

