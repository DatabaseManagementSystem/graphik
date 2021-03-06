<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: graphik.proto

namespace Api;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Refs is an array of refs
 *
 * Generated from protobuf message <code>api.Refs</code>
 */
class Refs extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .api.Ref refs = 1;</code>
     */
    private $refs;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $refs
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Graphik::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref refs = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getRefs()
    {
        return $this->refs;
    }

    /**
     * Generated from protobuf field <code>repeated .api.Ref refs = 1;</code>
     * @param \Api\Ref[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setRefs($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Api\Ref::class);
        $this->refs = $arr;

        return $this;
    }

}

