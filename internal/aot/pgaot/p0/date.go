package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateDate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	if l1 != 0 {
		if l0&int32(32768) != 0 {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
			v44 = v39 + int32(4799)
			v46 = base.I32_div_s(v44, int32(4))
			v49 = base.I32_div_s(v44, int32(-100))
			v52 = base.I32_div_s(v44, int32(400))
			v53 = v38 + v39*int32(365) + v46 + v49 + v52
			v55 = v53 + int32(1751940)
			v56 = int32(146097)
			v57 = base.I32_div_u_s(v55, v56)
			v58 = int32(3)
			v64 = int32(2)
			v69 = base.I32_div_u_s((v57*int32(1073595727)+v55)<<(uint(v64)%32)|v58, v56)
			v72 = v53 + v57*v58 + v69 + int32(1752000)
			v73 = int32(1461)
			v74 = base.I32_div_u_s(v72, v73)
			v77 = v74*int32(-1461) + v72
			v79 = v77 << (uint(v64) % 32)
			if base.Ui32(v73) <= base.Ui32(v79) {
				v85 = base.I32_rem_u_s(v77+int32(305), int32(365))
				v90 = v85
			} else {
				v89 = base.I32_rem_u_s(v77+int32(306), int32(366))
				v90 = v89
			}
			v92 = base.I32_div_u_s(v79, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v92 + v74<<(uint(int32(2))%32) - int32(4800)
			v100 = v90 + int32(123)
			v104 = int32(base.Ui32(v100*int32(2141)) >> (uint(int32(16)) % 32))
			v108 = base.I32_rem_u_s(v104+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v108 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v100 - int32(base.Ui32(v104*int32(7834))>>(uint(int32(8))%32))
		} else {
		}
		if l0&int32(2) == int32(0) {
			if l0&int32(8) == int32(0) {
				v143 = int32(14)
				if l0&v143 != v143 {
					return int32(0)
				} else {
					v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
					if v149&int32(3) != 0 {
						v159 = int32(0)
					} else {
						v154 = base.I32_rem_s(v149, int32(100))
						if v154 != 0 {
							v159 = int32(1)
						} else {
							v156 = base.I32_rem_s(v149, int32(400))
							v159 = base.B2i32(v156 == int32(0))
						}
					}
					v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
					if v147 <= v168 {
						return int32(0)
					} else {
						return int32(-2)
					}
				}
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
				if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
					v143 = int32(14)
					if l0&v143 != v143 {
						return int32(0)
					} else {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						if v149&int32(3) != 0 {
							v159 = int32(0)
						} else {
							v154 = base.I32_rem_s(v149, int32(100))
							if v154 != 0 {
								v159 = int32(1)
							} else {
								v156 = base.I32_rem_s(v149, int32(400))
								v159 = base.B2i32(v156 == int32(0))
							}
						}
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
						if v147 <= v168 {
							return int32(0)
						} else {
							return int32(-2)
						}
					}
				} else {
					return int32(-3)
				}
			}
		} else {
			v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
			if base.Ui32(int32(-12)) <= base.Ui32(v125-int32(13)) {
				if l0&int32(8) == int32(0) {
					v143 = int32(14)
					if l0&v143 != v143 {
						return int32(0)
					} else {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						if v149&int32(3) != 0 {
							v159 = int32(0)
						} else {
							v154 = base.I32_rem_s(v149, int32(100))
							if v154 != 0 {
								v159 = int32(1)
							} else {
								v156 = base.I32_rem_s(v149, int32(400))
								v159 = base.B2i32(v156 == int32(0))
							}
						}
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
						if v147 <= v168 {
							return int32(0)
						} else {
							return int32(-2)
						}
					}
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
						v143 = int32(14)
						if l0&v143 != v143 {
							return int32(0)
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							if v149&int32(3) != 0 {
								v159 = int32(0)
							} else {
								v154 = base.I32_rem_s(v149, int32(100))
								if v154 != 0 {
									v159 = int32(1)
								} else {
									v156 = base.I32_rem_s(v149, int32(400))
									v159 = base.B2i32(v156 == int32(0))
								}
							}
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
							if v147 <= v168 {
								return int32(0)
							} else {
								return int32(-2)
							}
						}
					} else {
						return int32(-3)
					}
				}
			} else {
				return int32(-3)
			}
		}
	} else {
		if l0&int32(4) == int32(0) {
			if l0&int32(32768) != 0 {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
				v44 = v39 + int32(4799)
				v46 = base.I32_div_s(v44, int32(4))
				v49 = base.I32_div_s(v44, int32(-100))
				v52 = base.I32_div_s(v44, int32(400))
				v53 = v38 + v39*int32(365) + v46 + v49 + v52
				v55 = v53 + int32(1751940)
				v56 = int32(146097)
				v57 = base.I32_div_u_s(v55, v56)
				v58 = int32(3)
				v64 = int32(2)
				v69 = base.I32_div_u_s((v57*int32(1073595727)+v55)<<(uint(v64)%32)|v58, v56)
				v72 = v53 + v57*v58 + v69 + int32(1752000)
				v73 = int32(1461)
				v74 = base.I32_div_u_s(v72, v73)
				v77 = v74*int32(-1461) + v72
				v79 = v77 << (uint(v64) % 32)
				if base.Ui32(v73) <= base.Ui32(v79) {
					v85 = base.I32_rem_u_s(v77+int32(305), int32(365))
					v90 = v85
				} else {
					v89 = base.I32_rem_u_s(v77+int32(306), int32(366))
					v90 = v89
				}
				v92 = base.I32_div_u_s(v79, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v92 + v74<<(uint(int32(2))%32) - int32(4800)
				v100 = v90 + int32(123)
				v104 = int32(base.Ui32(v100*int32(2141)) >> (uint(int32(16)) % 32))
				v108 = base.I32_rem_u_s(v104+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v108 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v100 - int32(base.Ui32(v104*int32(7834))>>(uint(int32(8))%32))
			} else {
			}
			if l0&int32(2) == int32(0) {
				if l0&int32(8) == int32(0) {
					v143 = int32(14)
					if l0&v143 != v143 {
						return int32(0)
					} else {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						if v149&int32(3) != 0 {
							v159 = int32(0)
						} else {
							v154 = base.I32_rem_s(v149, int32(100))
							if v154 != 0 {
								v159 = int32(1)
							} else {
								v156 = base.I32_rem_s(v149, int32(400))
								v159 = base.B2i32(v156 == int32(0))
							}
						}
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
						if v147 <= v168 {
							return int32(0)
						} else {
							return int32(-2)
						}
					}
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
						v143 = int32(14)
						if l0&v143 != v143 {
							return int32(0)
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							if v149&int32(3) != 0 {
								v159 = int32(0)
							} else {
								v154 = base.I32_rem_s(v149, int32(100))
								if v154 != 0 {
									v159 = int32(1)
								} else {
									v156 = base.I32_rem_s(v149, int32(400))
									v159 = base.B2i32(v156 == int32(0))
								}
							}
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
							if v147 <= v168 {
								return int32(0)
							} else {
								return int32(-2)
							}
						}
					} else {
						return int32(-3)
					}
				}
			} else {
				v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
				if base.Ui32(int32(-12)) <= base.Ui32(v125-int32(13)) {
					if l0&int32(8) == int32(0) {
						v143 = int32(14)
						if l0&v143 != v143 {
							return int32(0)
						} else {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							if v149&int32(3) != 0 {
								v159 = int32(0)
							} else {
								v154 = base.I32_rem_s(v149, int32(100))
								if v154 != 0 {
									v159 = int32(1)
								} else {
									v156 = base.I32_rem_s(v149, int32(400))
									v159 = base.B2i32(v156 == int32(0))
								}
							}
							v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
							if v147 <= v168 {
								return int32(0)
							} else {
								return int32(-2)
							}
						}
					} else {
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
							v143 = int32(14)
							if l0&v143 != v143 {
								return int32(0)
							} else {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v149&int32(3) != 0 {
									v159 = int32(0)
								} else {
									v154 = base.I32_rem_s(v149, int32(100))
									if v154 != 0 {
										v159 = int32(1)
									} else {
										v156 = base.I32_rem_s(v149, int32(400))
										v159 = base.B2i32(v156 == int32(0))
									}
								}
								v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
								if v147 <= v168 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							return int32(-3)
						}
					}
				} else {
					return int32(-3)
				}
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
			if l3 != 0 {
				if int32(0) < v10 {
					v33 = int32(1) - v10
					*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v33
					if l0&int32(32768) != 0 {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
						v44 = v39 + int32(4799)
						v46 = base.I32_div_s(v44, int32(4))
						v49 = base.I32_div_s(v44, int32(-100))
						v52 = base.I32_div_s(v44, int32(400))
						v53 = v38 + v39*int32(365) + v46 + v49 + v52
						v55 = v53 + int32(1751940)
						v56 = int32(146097)
						v57 = base.I32_div_u_s(v55, v56)
						v58 = int32(3)
						v64 = int32(2)
						v69 = base.I32_div_u_s((v57*int32(1073595727)+v55)<<(uint(v64)%32)|v58, v56)
						v72 = v53 + v57*v58 + v69 + int32(1752000)
						v73 = int32(1461)
						v74 = base.I32_div_u_s(v72, v73)
						v77 = v74*int32(-1461) + v72
						v79 = v77 << (uint(v64) % 32)
						if base.Ui32(v73) <= base.Ui32(v79) {
							v85 = base.I32_rem_u_s(v77+int32(305), int32(365))
							v90 = v85
						} else {
							v89 = base.I32_rem_u_s(v77+int32(306), int32(366))
							v90 = v89
						}
						v92 = base.I32_div_u_s(v79, int32(1461))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v92 + v74<<(uint(int32(2))%32) - int32(4800)
						v100 = v90 + int32(123)
						v104 = int32(base.Ui32(v100*int32(2141)) >> (uint(int32(16)) % 32))
						v108 = base.I32_rem_u_s(v104+int32(10), int32(12))
						*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v108 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v100 - int32(base.Ui32(v104*int32(7834))>>(uint(int32(8))%32))
					} else {
					}
					if l0&int32(2) == int32(0) {
						if l0&int32(8) == int32(0) {
							v143 = int32(14)
							if l0&v143 != v143 {
								return int32(0)
							} else {
								v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
								if v149&int32(3) != 0 {
									v159 = int32(0)
								} else {
									v154 = base.I32_rem_s(v149, int32(100))
									if v154 != 0 {
										v159 = int32(1)
									} else {
										v156 = base.I32_rem_s(v149, int32(400))
										v159 = base.B2i32(v156 == int32(0))
									}
								}
								v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
								if v147 <= v168 {
									return int32(0)
								} else {
									return int32(-2)
								}
							}
						} else {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
							if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
								v143 = int32(14)
								if l0&v143 != v143 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v149&int32(3) != 0 {
										v159 = int32(0)
									} else {
										v154 = base.I32_rem_s(v149, int32(100))
										if v154 != 0 {
											v159 = int32(1)
										} else {
											v156 = base.I32_rem_s(v149, int32(400))
											v159 = base.B2i32(v156 == int32(0))
										}
									}
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
									if v147 <= v168 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								return int32(-3)
							}
						}
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
						if base.Ui32(int32(-12)) <= base.Ui32(v125-int32(13)) {
							if l0&int32(8) == int32(0) {
								v143 = int32(14)
								if l0&v143 != v143 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v149&int32(3) != 0 {
										v159 = int32(0)
									} else {
										v154 = base.I32_rem_s(v149, int32(100))
										if v154 != 0 {
											v159 = int32(1)
										} else {
											v156 = base.I32_rem_s(v149, int32(400))
											v159 = base.B2i32(v156 == int32(0))
										}
									}
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
									if v147 <= v168 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
									v143 = int32(14)
									if l0&v143 != v143 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v149&int32(3) != 0 {
											v159 = int32(0)
										} else {
											v154 = base.I32_rem_s(v149, int32(100))
											if v154 != 0 {
												v159 = int32(1)
											} else {
												v156 = base.I32_rem_s(v149, int32(400))
												v159 = base.B2i32(v156 == int32(0))
											}
										}
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
										if v147 <= v168 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									return int32(-3)
								}
							}
						} else {
							return int32(-3)
						}
					}
				} else {
					return int32(-2)
				}
			} else {
				if l2 != 0 {
					if v10 < int32(0) {
						return int32(-2)
					} else {
						if base.Ui32(v10) <= base.Ui32(int32(69)) {
							v33 = v10 + int32(2000)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v33
						} else {
							if base.Ui32(int32(99)) < base.Ui32(v10) {
							} else {
								v33 = v10 + int32(1900)
								*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v33
							}
						}
						if l0&int32(32768) != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							v44 = v39 + int32(4799)
							v46 = base.I32_div_s(v44, int32(4))
							v49 = base.I32_div_s(v44, int32(-100))
							v52 = base.I32_div_s(v44, int32(400))
							v53 = v38 + v39*int32(365) + v46 + v49 + v52
							v55 = v53 + int32(1751940)
							v56 = int32(146097)
							v57 = base.I32_div_u_s(v55, v56)
							v58 = int32(3)
							v64 = int32(2)
							v69 = base.I32_div_u_s((v57*int32(1073595727)+v55)<<(uint(v64)%32)|v58, v56)
							v72 = v53 + v57*v58 + v69 + int32(1752000)
							v73 = int32(1461)
							v74 = base.I32_div_u_s(v72, v73)
							v77 = v74*int32(-1461) + v72
							v79 = v77 << (uint(v64) % 32)
							if base.Ui32(v73) <= base.Ui32(v79) {
								v85 = base.I32_rem_u_s(v77+int32(305), int32(365))
								v90 = v85
							} else {
								v89 = base.I32_rem_u_s(v77+int32(306), int32(366))
								v90 = v89
							}
							v92 = base.I32_div_u_s(v79, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v92 + v74<<(uint(int32(2))%32) - int32(4800)
							v100 = v90 + int32(123)
							v104 = int32(base.Ui32(v100*int32(2141)) >> (uint(int32(16)) % 32))
							v108 = base.I32_rem_u_s(v104+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v108 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v100 - int32(base.Ui32(v104*int32(7834))>>(uint(int32(8))%32))
						} else {
						}
						if l0&int32(2) == int32(0) {
							if l0&int32(8) == int32(0) {
								v143 = int32(14)
								if l0&v143 != v143 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v149&int32(3) != 0 {
										v159 = int32(0)
									} else {
										v154 = base.I32_rem_s(v149, int32(100))
										if v154 != 0 {
											v159 = int32(1)
										} else {
											v156 = base.I32_rem_s(v149, int32(400))
											v159 = base.B2i32(v156 == int32(0))
										}
									}
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
									if v147 <= v168 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
									v143 = int32(14)
									if l0&v143 != v143 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v149&int32(3) != 0 {
											v159 = int32(0)
										} else {
											v154 = base.I32_rem_s(v149, int32(100))
											if v154 != 0 {
												v159 = int32(1)
											} else {
												v156 = base.I32_rem_s(v149, int32(400))
												v159 = base.B2i32(v156 == int32(0))
											}
										}
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
										if v147 <= v168 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									return int32(-3)
								}
							}
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if base.Ui32(int32(-12)) <= base.Ui32(v125-int32(13)) {
								if l0&int32(8) == int32(0) {
									v143 = int32(14)
									if l0&v143 != v143 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v149&int32(3) != 0 {
											v159 = int32(0)
										} else {
											v154 = base.I32_rem_s(v149, int32(100))
											if v154 != 0 {
												v159 = int32(1)
											} else {
												v156 = base.I32_rem_s(v149, int32(400))
												v159 = base.B2i32(v156 == int32(0))
											}
										}
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
										if v147 <= v168 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
										v143 = int32(14)
										if l0&v143 != v143 {
											return int32(0)
										} else {
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
											if v149&int32(3) != 0 {
												v159 = int32(0)
											} else {
												v154 = base.I32_rem_s(v149, int32(100))
												if v154 != 0 {
													v159 = int32(1)
												} else {
													v156 = base.I32_rem_s(v149, int32(400))
													v159 = base.B2i32(v156 == int32(0))
												}
											}
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
											if v147 <= v168 {
												return int32(0)
											} else {
												return int32(-2)
											}
										}
									} else {
										return int32(-3)
									}
								}
							} else {
								return int32(-3)
							}
						}
					}
				} else {
					if int32(0) < v10 {
						if l0&int32(32768) != 0 {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
							v44 = v39 + int32(4799)
							v46 = base.I32_div_s(v44, int32(4))
							v49 = base.I32_div_s(v44, int32(-100))
							v52 = base.I32_div_s(v44, int32(400))
							v53 = v38 + v39*int32(365) + v46 + v49 + v52
							v55 = v53 + int32(1751940)
							v56 = int32(146097)
							v57 = base.I32_div_u_s(v55, v56)
							v58 = int32(3)
							v64 = int32(2)
							v69 = base.I32_div_u_s((v57*int32(1073595727)+v55)<<(uint(v64)%32)|v58, v56)
							v72 = v53 + v57*v58 + v69 + int32(1752000)
							v73 = int32(1461)
							v74 = base.I32_div_u_s(v72, v73)
							v77 = v74*int32(-1461) + v72
							v79 = v77 << (uint(v64) % 32)
							if base.Ui32(v73) <= base.Ui32(v79) {
								v85 = base.I32_rem_u_s(v77+int32(305), int32(365))
								v90 = v85
							} else {
								v89 = base.I32_rem_u_s(v77+int32(306), int32(366))
								v90 = v89
							}
							v92 = base.I32_div_u_s(v79, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+20)) = v92 + v74<<(uint(int32(2))%32) - int32(4800)
							v100 = v90 + int32(123)
							v104 = int32(base.Ui32(v100*int32(2141)) >> (uint(int32(16)) % 32))
							v108 = base.I32_rem_u_s(v104+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v108 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v100 - int32(base.Ui32(v104*int32(7834))>>(uint(int32(8))%32))
						} else {
						}
						if l0&int32(2) == int32(0) {
							if l0&int32(8) == int32(0) {
								v143 = int32(14)
								if l0&v143 != v143 {
									return int32(0)
								} else {
									v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
									if v149&int32(3) != 0 {
										v159 = int32(0)
									} else {
										v154 = base.I32_rem_s(v149, int32(100))
										if v154 != 0 {
											v159 = int32(1)
										} else {
											v156 = base.I32_rem_s(v149, int32(400))
											v159 = base.B2i32(v156 == int32(0))
										}
									}
									v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
									if v147 <= v168 {
										return int32(0)
									} else {
										return int32(-2)
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
									v143 = int32(14)
									if l0&v143 != v143 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v149&int32(3) != 0 {
											v159 = int32(0)
										} else {
											v154 = base.I32_rem_s(v149, int32(100))
											if v154 != 0 {
												v159 = int32(1)
											} else {
												v156 = base.I32_rem_s(v149, int32(400))
												v159 = base.B2i32(v156 == int32(0))
											}
										}
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
										if v147 <= v168 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									return int32(-3)
								}
							}
						} else {
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
							if base.Ui32(int32(-12)) <= base.Ui32(v125-int32(13)) {
								if l0&int32(8) == int32(0) {
									v143 = int32(14)
									if l0&v143 != v143 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
										if v149&int32(3) != 0 {
											v159 = int32(0)
										} else {
											v154 = base.I32_rem_s(v149, int32(100))
											if v154 != 0 {
												v159 = int32(1)
											} else {
												v156 = base.I32_rem_s(v149, int32(400))
												v159 = base.B2i32(v156 == int32(0))
											}
										}
										v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
										if v147 <= v168 {
											return int32(0)
										} else {
											return int32(-2)
										}
									}
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									if base.Ui32(int32(-31)) <= base.Ui32(v136-int32(32)) {
										v143 = int32(14)
										if l0&v143 != v143 {
											return int32(0)
										} else {
											v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
											if v149&int32(3) != 0 {
												v159 = int32(0)
											} else {
												v154 = base.I32_rem_s(v149, int32(100))
												if v154 != 0 {
													v159 = int32(1)
												} else {
													v156 = base.I32_rem_s(v149, int32(400))
													v159 = base.B2i32(v156 == int32(0))
												}
											}
											v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v159*int32(52)+v162<<(uint(int32(2))%32))+uint32(_consts[786])))
											if v147 <= v168 {
												return int32(0)
											} else {
												return int32(-2)
											}
										}
									} else {
										return int32(-3)
									}
								}
							} else {
								return int32(-3)
							}
						}
					} else {
						return int32(-2)
					}
				}
			}
		}
	}
}
func F_date_cmp_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v8 int64
	_ = v8
	var v15 int64
	_ = v15
	var v24 int64
	_ = v24
	var v33 int32
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 == int32(-2147483648) {
		v8 = int64(-9223372036854775807 - 1)
		return base.B2i32(v4 < v8) - base.B2i32(v8 < v4)
	} else {
		if v5 == int32(2147483647) {
			v15 = int64(9223372036854775807)
			return base.B2i32(v4 < v15) - base.B2i32(v15 < v4)
		} else {
			if v5 <= int32(106751982) {
				v24 = base.I64_extend_i32_s(v5) * int64(86400000000)
				return base.B2i32(v4 < v24) - base.B2i32(v24 < v4)
			} else {
				if v4 == int64(9223372036854775807) {
					v33 = int32(-1)
				} else {
					v33 = int32(1)
				}
				return v33
			}
		}
	}
}
func F_date_cmp_timestamptz_internal(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v119 int64
	_ = v119
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v142 int32
	_ = v142
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l0 == int32(-2147483648) {
		v11 = int64(-9223372036854775807 - 1)
		v142 = base.B2i32(l1 < v11) - base.B2i32(v11 < l1)
	} else {
		if l0 == int32(2147483647) {
			v137 = int64(9223372036854775807)
			v142 = base.B2i32(l1 < v137) - base.B2i32(v137 < l1)
		} else {
			if l0 <= int32(106751982) {
				v31 = l0 + int32(2483589)
				v32 = int32(146097)
				v33 = base.I32_div_u_s(v31, v32)
				v34 = int32(3)
				v40 = int32(2)
				v45 = base.I32_div_u_s((v33*int32(1073595727)+v31)<<(uint(v40)%32)|v34, v32)
				v48 = l0 + int32(2451545) + v33*v34 + v45 + int32(32104)
				v49 = int32(1461)
				v50 = base.I32_div_u_s(v48, v49)
				v53 = v50*int32(-1461) + v48
				v55 = v53 << (uint(v40) % 32)
				if base.Ui32(v49) <= base.Ui32(v55) {
					v61 = base.I32_rem_u_s(v53+int32(305), int32(365))
					v66 = v61
				} else {
					v65 = base.I32_rem_u_s(v53+int32(306), int32(366))
					v66 = v65
				}
				v68 = base.I32_div_u_s(v55, int32(1461))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v68 + v50<<(uint(int32(2))%32) - int32(4800)
				v76 = v66 + int32(123)
				v80 = int32(base.Ui32(v76*int32(2141)) >> (uint(int32(16)) % 32))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v76 - int32(base.Ui32(v80*int32(7834))>>(uint(int32(8))%32))
				v90 = base.I32_rem_u_s(v80+int32(10), int32(12))
				*(*int32)(unsafe.Add(mBase, uint32(v7+int32(20)))) = v90 + int32(1)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
				v101 = *(*int32)(unsafe.Add(mBase, _consts[774]))
				v103 = m.G0
				v104 = int32(16)
				v105 = v103 - v104
				m.G0 = v105
				v109 = F_DetermineTimeZoneOffsetInternal(m, v7+int32(4), v101, v105+int32(8))
				mBase = m.M
				m.G0 = v105 + v104
				v119 = base.I64_extend_i32_s(v109)*int64(1000000) + base.I64_extend_i32_s(l0)*int64(86400000000)
				if base.Ui64(v119+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v137 = v119
					v142 = base.B2i32(l1 < v137) - base.B2i32(v137 < l1)
				} else {
					if v119 < int64(-211813488000000000) {
						if l1 == int64(-9223372036854775807-1) {
							v136 = int32(1)
						} else {
							v136 = int32(-1)
						}
						v142 = v136
					} else {
						if l1 == int64(9223372036854775807) {
							v131 = int32(-1)
						} else {
							v131 = int32(1)
						}
						v142 = v131
					}
				}
			} else {
				if l1 == int64(9223372036854775807) {
					v131 = int32(-1)
				} else {
					v131 = int32(1)
				}
				v142 = v131
			}
		}
	}
	m.G0 = v7 + int32(48)
	return v142
}
func F_date_finite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.B2i32(base.Ui32(v2+int32(2147483647)) < base.Ui32(int32(-2)))
}
func F_date_ge_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
		} else {
			if int32(106751982) < v6 {
				if v4 == int64(9223372036854775807) {
					v25 = int32(-1)
				} else {
					v25 = int32(1)
				}
				v26 = v25
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
			}
		}
	}
	return int32(base.Ui32(v26^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_date_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v3 < v2)
}
func F_date_mii(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v4-int32(2147483647)) < base.Ui32(int32(2)) {
		v37 = v4
		return v37
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = v4 - v9
		if int32(0) <= v9 {
			if v10 <= v4 {
				if base.Ui32(v10+int32(2451545)) < base.Ui32(int32(2147483494)) {
					v37 = v10
					return v37
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(402633), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498806), int32(609), int32(320150))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(402633), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498806), int32(609), int32(320150))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			if v10 < v4 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(402633), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498806), int32(609), int32(320150))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(v10+int32(2451545)) < base.Ui32(int32(2147483494)) {
					v37 = v10
					return v37
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(402633), int32(0))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(498806), int32(609), int32(320150))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
