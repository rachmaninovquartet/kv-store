#!/bin/bash

# Test All Endpoints Script
# This script tests all endpoints in the test client and prints the results

echo "🧪 Testing All Test Client Endpoints"
echo "=================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test client base URL
TEST_CLIENT_URL="http://localhost:8002"

echo -e "${BLUE}1. Testing Deletion Workflow${NC}"
echo "--------------------------------"
response=$(curl -s "http://localhost:8002/test_deletion")
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Success${NC}"
    echo "$response" | python3 -m json.tool 2>/dev/null || echo "$response"
else
    echo -e "${RED}❌ Failed${NC}"
    echo "$response"
fi
echo ""

echo -e "${BLUE}2. Testing Overwrite Workflow${NC}"
echo "----------------------------------"
response=$(curl -s "http://localhost:8002/test_overwrite")
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Success${NC}"
    echo "$response" | python3 -m json.tool 2>/dev/null || echo "$response"
else
    echo -e "${RED}❌ Failed${NC}"
    echo "$response"
fi
echo ""


echo "=================================="
echo -e "${YELLOW}🎉 All tests completed!${NC}"
echo ""
