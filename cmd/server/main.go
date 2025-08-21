<?xml version="1.0" encoding="utf-8"?>
<wsdl:definitions xmlns:soap12="http://schemas.xmlsoap.org/wsdl/soap12/" xmlns:s="http://www.w3.org/2001/XMLSchema" xmlns:tm="http://microsoft.com/wsdl/mime/textMatching/" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:http="http://schemas.xmlsoap.org/wsdl/http/" xmlns:tns="http://tempuri.org/" xmlns:mime="http://schemas.xmlsoap.org/wsdl/mime/" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" targetNamespace="http://tempuri.org/" xmlns:wsdl="http://schemas.xmlsoap.org/wsdl/">
  <wsdl:types>
    <s:schema elementFormDefault="qualified" targetNamespace="http://tempuri.org/">
      <s:element name="CheckCardByLaser">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="PID" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="FirstName" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="LastName" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="BirthDay" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="Laser" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="CheckCardByLaserResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="CheckCardByLaserResult" type="tns:CardStatusOut" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:complexType name="CardStatusOut">
        <s:complexContent mixed="false">
          <s:extension base="tns:ResponseOut">
            <s:sequence>
              <s:element minOccurs="1" maxOccurs="1" name="Code" type="s:unsignedShort" />
              <s:element minOccurs="0" maxOccurs="1" name="Desc" type="s:string" />
            </s:sequence>
          </s:extension>
        </s:complexContent>
      </s:complexType>
      <s:complexType name="ResponseOut">
        <s:sequence>
          <s:element minOccurs="1" maxOccurs="1" name="IsError" type="s:boolean" />
          <s:element minOccurs="0" maxOccurs="1" name="ErrorMessage" type="s:string" />
        </s:sequence>
      </s:complexType>
      <s:element name="CheckCardByCID">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="ChipNo" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="pid" type="s:string" />
            <s:element minOccurs="0" maxOccurs="1" name="bp1no" type="s:string" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="CheckCardByCIDResponse">
        <s:complexType>
          <s:sequence>
            <s:element minOccurs="0" maxOccurs="1" name="CheckCardByCIDResult" type="tns:CardStatusOut" />
          </s:sequence>
        </s:complexType>
      </s:element>
      <s:element name="CardStatusOut" nillable="true" type="tns:CardStatusOut" />
    </s:schema>
  </wsdl:types>
  <wsdl:message name="CheckCardByLaserSoapIn">
    <wsdl:part name="parameters" element="tns:CheckCardByLaser" />
  </wsdl:message>
  <wsdl:message name="CheckCardByLaserSoapOut">
    <wsdl:part name="parameters" element="tns:CheckCardByLaserResponse" />
  </wsdl:message>
  <wsdl:message name="CheckCardByCIDSoapIn">
    <wsdl:part name="parameters" element="tns:CheckCardByCID" />
  </wsdl:message>
  <wsdl:message name="CheckCardByCIDSoapOut">
    <wsdl:part name="parameters" element="tns:CheckCardByCIDResponse" />
  </wsdl:message>
  <wsdl:message name="CheckCardByLaserHttpGetIn">
    <wsdl:part name="PID" type="s:string" />
    <wsdl:part name="FirstName" type="s:string" />
    <wsdl:part name="LastName" type="s:string" />
    <wsdl:part name="BirthDay" type="s:string" />
    <wsdl:part name="Laser" type="s:string" />
  </wsdl:message>
  <wsdl:message name="CheckCardByLaserHttpGetOut">
    <wsdl:part name="Body" element="tns:CardStatusOut" />
  </wsdl:message>
  <wsdl:message name="CheckCardByCIDHttpGetIn">
    <wsdl:part name="ChipNo" type="s:string" />
    <wsdl:part name="pid" type="s:string" />
    <wsdl:part name="bp1no" type="s:string" />
  </wsdl:message>
  <wsdl:message name="CheckCardByCIDHttpGetOut">
    <wsdl:part name="Body" element="tns:CardStatusOut" />
  </wsdl:message>
  <wsdl:portType name="CheckCardBankServiceSoap">
    <wsdl:operation name="CheckCardByLaser">
      <wsdl:input message="tns:CheckCardByLaserSoapIn" />
      <wsdl:output message="tns:CheckCardByLaserSoapOut" />
    </wsdl:operation>
    <wsdl:operation name="CheckCardByCID">
      <wsdl:input message="tns:CheckCardByCIDSoapIn" />
      <wsdl:output message="tns:CheckCardByCIDSoapOut" />
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:portType name="CheckCardBankServiceHttpGet">
    <wsdl:operation name="CheckCardByLaser">
      <wsdl:input message="tns:CheckCardByLaserHttpGetIn" />
      <wsdl:output message="tns:CheckCardByLaserHttpGetOut" />
    </wsdl:operation>
    <wsdl:operation name="CheckCardByCID">
      <wsdl:input message="tns:CheckCardByCIDHttpGetIn" />
      <wsdl:output message="tns:CheckCardByCIDHttpGetOut" />
    </wsdl:operation>
  </wsdl:portType>
  <wsdl:binding name="CheckCardBankServiceSoap" type="tns:CheckCardBankServiceSoap">
    <soap:binding transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="CheckCardByLaser">
      <soap:operation soapAction="http://tempuri.org/CheckCardByLaser" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="CheckCardByCID">
      <soap:operation soapAction="http://tempuri.org/CheckCardByCID" style="document" />
      <wsdl:input>
        <soap:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:binding name="CheckCardBankServiceSoap12" type="tns:CheckCardBankServiceSoap">
    <soap12:binding transport="http://schemas.xmlsoap.org/soap/http" />
    <wsdl:operation name="CheckCardByLaser">
      <soap12:operation soapAction="http://tempuri.org/CheckCardByLaser" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="CheckCardByCID">
      <soap12:operation soapAction="http://tempuri.org/CheckCardByCID" style="document" />
      <wsdl:input>
        <soap12:body use="literal" />
      </wsdl:input>
      <wsdl:output>
        <soap12:body use="literal" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:binding name="CheckCardBankServiceHttpGet" type="tns:CheckCardBankServiceHttpGet">
    <http:binding verb="GET" />
    <wsdl:operation name="CheckCardByLaser">
      <http:operation location="/CheckCardByLaser" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
    <wsdl:operation name="CheckCardByCID">
      <http:operation location="/CheckCardByCID" />
      <wsdl:input>
        <http:urlEncoded />
      </wsdl:input>
      <wsdl:output>
        <mime:mimeXml part="Body" />
      </wsdl:output>
    </wsdl:operation>
  </wsdl:binding>
  <wsdl:service name="CheckCardBankService">
    <wsdl:port name="CheckCardBankServiceSoap" binding="tns:CheckCardBankServiceSoap">
      <soap:address location="https://idcard.bora.dopa.go.th/CheckCardStatus/CheckCardService.asmx" />
    </wsdl:port>
    <wsdl:port name="CheckCardBankServiceSoap12" binding="tns:CheckCardBankServiceSoap12">
      <soap12:address location="https://idcard.bora.dopa.go.th/CheckCardStatus/CheckCardService.asmx" />
    </wsdl:port>
    <wsdl:port name="CheckCardBankServiceHttpGet" binding="tns:CheckCardBankServiceHttpGet">
      <http:address location="https://idcard.bora.dopa.go.th/CheckCardStatus/CheckCardService.asmx" />
    </wsdl:port>
  </wsdl:service>
</wsdl:definitions>
