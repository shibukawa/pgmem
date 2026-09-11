package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbValueToJsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v9) < base.Ui32(int32(4)) {
		v15 = F_palloc(m, int32(32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v15)+28)) = uint16(v19)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
			v24 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v24
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v24)
			*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(16)
			v31 = F_palloc(m, int32(20))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v31
				v37 = F_pushJsonbValue(m, v7+int32(8), int32(3), l0)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
					if v40 == int32(0) {
						F_initStringInfo(m, v7+int32(16))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_enlargeStringInfo(m, v7+int32(16), int32(4))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
								v89 = v87 + int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								v93 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
								F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
									v154 = v102
									m.G0 = v7 + int32(32)
									return v154
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
						switch v44 - int32(16) {
						case 0:
							F_appendElement(m, v40, v39)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_initStringInfo(m, v7+int32(16))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v7+int32(16), int32(4))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										v89 = v87 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v93 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
										F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
											v154 = v102
											m.G0 = v7 + int32(32)
											return v154
										}
									}
								}
							}
						case 1:
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v47 + int32(1)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
							v54 = v51 + v47*int32(44)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v55
							v57 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v54)+28)) = v57
							v59 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
							*(*int64)(unsafe.Add(mBase, uint32(v54)+20)) = v59
							F_initStringInfo(m, v7+int32(16))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v7+int32(16), int32(4))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									v89 = v87 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v93 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
									F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
										v154 = v102
										m.G0 = v7 + int32(32)
										return v154
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(345964), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(466844), int32(720), int32(216117))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
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
	} else {
		switch v9 - int32(16) {
		case 0, 1:
			F_initStringInfo(m, v7+int32(16))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v7+int32(16), int32(4))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v118 = v116 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v118
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
					v122 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v120+v118))) = uint8(v122)
					F_convertJsonbValue(m, v7+int32(16), v7+int32(12), l0, v122)
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v131))) = v132 << (uint(int32(2)) % 32)
						v154 = v131
						m.G0 = v7 + int32(32)
						return v154
					}
				}
			}
		default:
			v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v139 = F_palloc(m, v136+int32(4))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int32(0)
			} else {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v139))) = v141<<(uint(int32(2))%32) + int32(16)
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v150 != 0 {
					v151 = F__emscripten_memcpy_bulkmem(m, v139+int32(4), v149, v150)
					mBase = m.M
				} else {
				}
				v154 = v139
				m.G0 = v7 + int32(32)
				return v154
			}
		case 16:
			v15 = F_palloc(m, int32(32))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+28)) = uint16(v19)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
				v24 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v24
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v24)
				*(*int64)(unsafe.Add(mBase, uint32(v15))) = int64(16)
				v31 = F_palloc(m, int32(20))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v31
					v37 = F_pushJsonbValue(m, v7+int32(8), int32(3), l0)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v40
						if v40 == int32(0) {
							F_initStringInfo(m, v7+int32(16))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v7+int32(16), int32(4))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									v89 = v87 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v93 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
									F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
										v154 = v102
										m.G0 = v7 + int32(32)
										return v154
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
							switch v44 - int32(16) {
							case 0:
								F_appendElement(m, v40, v39)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									F_initStringInfo(m, v7+int32(16))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_enlargeStringInfo(m, v7+int32(16), int32(4))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											v89 = v87 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v93 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
											F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int32(0)
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
												v154 = v102
												m.G0 = v7 + int32(32)
												return v154
											}
										}
									}
								}
							case 1:
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v47 + int32(1)
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
								v54 = v51 + v47*int32(44)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v54)+36)) = v55
								v57 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v54)+28)) = v57
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
								*(*int64)(unsafe.Add(mBase, uint32(v54)+20)) = v59
								F_initStringInfo(m, v7+int32(16))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v7+int32(16), int32(4))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										v89 = v87 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v89
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v93 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v91+v89))) = uint8(v93)
										F_convertJsonbValue(m, v7+int32(16), v7+int32(12), v39, v93)
										mBase = m.M
										v101 = m.ExcPending
										if v101 != 0 {
											return int32(0)
										} else {
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v102))) = v103 << (uint(int32(2)) % 32)
											v154 = v102
											m.G0 = v7 + int32(32)
											return v154
										}
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(345964), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(466844), int32(720), int32(216117))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
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
	}
}
func F_jsonb_array_elements(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_elements_worker_jsonb(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_build_array_noargs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v7
	v13 = F_pushJsonbValue(m, v5, int32(4), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v19 = F_pushJsonbValue(m, v5, int32(5), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v19
			v22 = F_JsonbValueToJsonb(m, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v5 + int32(16)
				return v22
			}
		}
	}
}
func F_jsonb_build_array_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v15
	v21 = F_pushJsonbValue(m, v13, int32(4), v6)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v21
	if int32(0) < l0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = v6
	goto L6
L4:
	;
	goto L5
L5:
	;
	v74 = F_pushJsonbValue(m, v13, int32(5), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v36))))
	if v42&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v45 = l4
	goto L10
L9:
	;
	v45 = int32(0)
	goto L10
L10:
	;
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = v36 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1+v49)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l3+v49)))
	F_add_jsonb(m, v51, (l4^int32(1))&v42, v13, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v60 = v36 + int32(1)
	if v60 != l0 {
		v36 = v60
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	goto L7
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v74
	v77 = F_JsonbValueToJsonb(m, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v13 + int32(16)
	return v77
}
func F_jsonb_build_object_noargs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v7
	v13 = F_pushJsonbValue(m, v5, int32(6), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v19 = F_pushJsonbValue(m, v5, int32(7), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v19
			v22 = F_JsonbValueToJsonb(m, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				m.G0 = v5 + int32(16)
				return v22
			}
		}
	}
}
func F_jsonb_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if (v17^v18)&int32(536870912) == int32(0) {
				v26 = F_JsonbIteratorInit(m, v15+int32(4))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v26
					v31 = F_JsonbIteratorInit(m, v10+int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v31
						v38 = F_JsonbDeepContains(m, v7+int32(12), v7+int32(8))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = v38
							m.G0 = v7 + int32(16)
							return v40
						}
					}
				}
			} else {
				v40 = int32(0)
				m.G0 = v7 + int32(16)
				return v40
			}
		}
	}
}
func F_jsonb_delete_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v29 < int32(2) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L80
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v32&int32(268435456) != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L76
	}
L8:
	;
	if v32&int32(268435455) == int32(0) {
		v288 = v20
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v17 + int32(48)
	return v288
L10:
	;
	F_deconstruct_array_builtin(m, v25, int32(25), v17+int32(44), v17+int32(40), v17+int32(36))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v48 == int32(0) {
		v288 = v20
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v53 = F_JsonbIteratorInit(m, v20+int32(4))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v53
	v61 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v284 = F_JsonbValueToJsonb(m, v275)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L75
	}
L15:
	;
	if v61 == int32(0) {
		v275 = v2
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v69 = v61
	v70 = v2
	goto L17
L17:
	;
	v80 = base.B2i32(v69 != int32(1))
	if v80&base.B2i32(v69 != int32(3)) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v275 = v70
	goto L14
L19:
	;
	if v80 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L20:
	;
	if base.Ui32(v69) < base.Ui32(int32(4)) {
		goto L63
	} else {
		goto L64
	}
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v84 != int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if v87 <= int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v97 = int32(0)
	goto L24
L24:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+v94))))
	if v110 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L20
L26:
	;
	v221 = v97 + int32(1)
	if v221 != v87 {
		v97 = v221
		goto L24
	} else {
		goto L62
	}
L27:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93+v97<<(uint(int32(2))%32))))
	v115 = int32(1)
	v116 = v114 + v115
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v119 = v117 & v115
	if v117 == v115 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v147 != v92 {
		goto L26
	} else {
		goto L39
	}
L29:
	;
	v122 = int32(4)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v124&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v137 = int32(1)
	if v119 != 0 {
		v147 = int32(base.Ui32(v117)>>(uint(v137)%32)) - v137
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v133 = v122
	goto L34
L33:
	;
	v133 = base.B2i32(v124 == int32(18)) << (uint(v122) % 32)
	goto L34
L34:
	;
	if v124 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v136 = v122
	goto L37
L36:
	;
	v136 = v133
	goto L37
L37:
	;
	v147 = v136
	goto L28
L38:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v147 = int32(base.Ui32(v141)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v151 = v116
	goto L42
L41:
	;
	v151 = v114 + int32(4)
	goto L42
L42:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v92) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v213 == int32(0) {
		goto L19
	} else {
		goto L61
	}
L44:
	;
	v213 = int32(0)
	goto L43
L45:
	;
	v187 = v182
	v188 = v183
	v189 = v184
	goto L55
L46:
	;
	if (v151|v91)&int32(3) != 0 {
		v182 = v151
		v183 = v91
		v184 = v92
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v175 = v151
	v176 = v91
	v177 = v92
	goto L48
L48:
	;
	if v177 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L49:
	;
	v159 = v151
	v160 = v91
	v161 = v92
	goto L50
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v164 != v165 {
		v182 = v159
		v183 = v160
		v184 = v161
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v175 = v170
	v176 = v168
	v177 = v172
	goto L48
L52:
	;
	v167 = int32(4)
	v168 = v160 + v167
	v170 = v159 + v167
	v172 = v161 - v167
	if base.Ui32(int32(3)) < base.Ui32(v172) {
		v159 = v170
		v160 = v168
		v161 = v172
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v182 = v175
	v183 = v176
	v184 = v177
	goto L45
L55:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == v193 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v213 = v192 - v193
	goto L43
L57:
	;
	v195 = int32(1)
	v200 = v189 - v195
	if v200 != 0 {
		v187 = v187 + v195
		v188 = v188 + v195
		v189 = v200
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L44
L61:
	;
	goto L26
L62:
	;
	goto L25
L63:
	;
	v244 = v17 + int32(8)
	goto L65
L64:
	;
	v244 = int32(0)
	goto L65
L65:
	;
	v245 = F_pushJsonbValue(m, v17+int32(32), v69, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v252 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v252 != 0 {
		v69 = v252
		v70 = v245
		goto L17
	} else {
		goto L68
	}
L68:
	;
	v275 = v245
	goto L14
L69:
	;
	v261 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v268 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(1))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	if v268 != 0 {
		v69 = v268
		goto L17
	} else {
		goto L74
	}
L74:
	;
	goto L18
L75:
	;
	v288 = v284
	goto L9
L76:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(109071), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(464335), int32(4734), int32(22404))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(215944), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(464335), int32(4739), int32(22404))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_exec_setup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = F_palloc0(m, v9<<(uint(int32(3))%32)+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v16)
	v19 = v14 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v19 + v9<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v14
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v26 == v16 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(1402)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(1403)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(1404)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1405)
	return
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v36 = int32(0)
	goto L6
L6:
	;
	v40 = v36 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41)))
	v44 = F_exprType(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v46+v40))) = v44
	v50 = v36 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v50 < v51 {
		v36 = v50
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_jsonb_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v20 = l0 + int32(4)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = int32(base.Ui32(v22&int32(536870912)) >> (uint(int32(29)) % 32))
	if v26 != 0 {
		v41 = v6
		v42 = v6
		goto L8
	} else {
		goto L9
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v244
L2:
	;
	if l4 == int32(0) {
		v244 = l0
		goto L1
	} else {
		goto L72
	}
L3:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v231)
	v244 = int32(0)
	goto L1
L4:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v228)
	v244 = int32(0)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L12
	} else {
		goto L69
	}
L6:
	;
	if l4 != 0 {
		goto L61
	} else {
		goto L62
	}
L7:
	;
	v55 = int32(0)
	v61 = v26
	v62 = v51
	v63 = v20
	goto L16
L8:
	;
	v43 = int32(0)
	if base.B2i32(v42 == v43)&base.B2i32(l2 <= v43) != 0 {
		goto L2
	} else {
		goto L14
	}
L9:
	;
	if v22&int32(1342177280) == int32(1073741824) {
		v41 = int32(1)
		v42 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(0)
	if v33 < l2 {
		v51 = v33
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v37 = F_getIthJsonbValueFromContainer(m, v20, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v41 = v33
	v42 = v37
	goto L8
L14:
	;
	if l2 <= int32(0) {
		v197 = v42
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v51 = v41
	goto L7
L16:
	;
	if v61&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v164 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(int32(2))%32))))
	v73 = F_pg_detoast_datum_packed(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v62&int32(1) == int32(0) {
		goto L3
	} else {
		goto L41
	}
L22:
	;
	v75 = int32(1)
	v76 = v73 + v75
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v81 = v79 & v75
	if v81 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v82 = v76
	goto L25
L24:
	;
	v82 = v73 + int32(4)
	goto L25
L25:
	;
	if v79 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v81 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v101 = F_getKeyJsonValueFromContainer(m, v63, v82, v99, int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v164 = v101
	goto L18
L36:
	;
	v103 = int32(1)
	v108 = F_getKeyJsonValueFromContainer(m, v63, v82, int32(base.Ui32(v79)>>(uint(v103)%32))-v103, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L12
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v116 = F_getKeyJsonValueFromContainer(m, v63, v82, int32(base.Ui32(v110)>>(uint(int32(2))%32))-int32(4), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	v164 = v108
	goto L18
L40:
	;
	v164 = v116
	goto L18
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1+v55<<(uint(int32(2))%32))))
	v126 = F_text_to_cstring(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v134 = F_strtol(m, v126, v15+int32(12), int32(10))
	mBase = m.M
	goto L43
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v126 == v135 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	if v137 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v139 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v134 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v142&int32(1073741824) == int32(0) {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v156 = v134
	goto L49
L49:
	;
	v157 = F_getIthJsonbValueFromContainer(m, v63, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L12
	} else {
		goto L53
	}
L50:
	;
	if v134 == int32(-2147483648) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v150 = v142 & int32(268435455)
	if base.Ui32(v150) < base.Ui32(int32(0)-v134) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v156 = v134 + v150
	goto L49
L53:
	;
	v164 = v157
	goto L18
L54:
	;
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v167)
	v244 = int32(0)
	goto L1
L55:
	;
	goto L56
L56:
	;
	if v55 == l2-int32(1) {
		v197 = v164
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v173 == int32(18) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v187 = int32(base.Ui32(v177&int32(536870912)) >> (uint(int32(29)) % 32))
	v188 = int32(base.Ui32(v177&int32(1073741824)) >> (uint(int32(30)) % 32))
	v189 = v176
	goto L60
L59:
	;
	v187 = v171
	v188 = v171
	v189 = v63
	goto L60
L60:
	;
	v55 = v55 + int32(1)
	v61 = v187
	v62 = v188
	v63 = v189
	goto L16
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v204 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v212 = F_JsonbValueToJsonb(m, v197)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L68
	}
L64:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v207)
	v244 = int32(0)
	goto L1
L65:
	;
	goto L66
L66:
	;
	v210 = F_JsonbValueAsText(m, v197)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	v244 = v210
	goto L1
L68:
	;
	v244 = v212
	goto L1
L69:
	;
	F_errmsg_internal(m, int32(23885), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(464335), int32(1614), int32(89198))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v240 = F_JsonbToCString(m, int32(0), v20, int32(base.Ui32(v237)>>(uint(int32(2))%32)))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v242 = F_cstring_to_text(m, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v244 = v242
	goto L1
}
func F_jsonb_in_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_pushJsonbValue(m, l0, int32(5), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
		return int32(0)
	}
}
func F_jsonb_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v41 = int32(0)
							m.G0 = v8 + int32(32)
							return v41
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v41 = int32(0)
						m.G0 = v8 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(212149))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v35 = F_DirectFunctionCall1Coll(m, int32(1333), int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v37 {
							v41 = v35
							m.G0 = v8 + int32(32)
							return v41
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v8 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(212149))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
func F_jsonb_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v20 <= int32(0))
							}
						} else {
							return base.B2i32(v20 <= int32(0))
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v20 <= int32(0))
						}
					} else {
						return base.B2i32(v20 <= int32(0))
					}
				}
			}
		}
	}
}
func F_jsonb_path_match(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_match_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_subscript_assign(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+44)))
	if v17 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
		if v39 == int32(1) {
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v43 == int32(1) {
				v46 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v46)
				v49 = int32(16)
			} else {
				v49 = int32(17)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v49
			v55 = F_JsonbValueToJsonb(m, v13+int32(8))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v58 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
				v64 = v55
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v68 = v13 + int32(28)
				v69 = m.G0
				v71 = v69 - int32(16)
				m.G0 = v71
				*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(0)
				v75 = F_palloc0(m, v66)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					if v77 != int32(16) {
					} else {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
						if v80 != int32(1) {
						} else {
							v84 = v13 + int32(36)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							*(*int64)(unsafe.Add(mBase, uint32(v68))) = v86
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v88
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v84))) = v90
						}
					}
					v96 = F_JsonbIteratorInit(m, v64+int32(4))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v96
						v105 = F_setPath(m, v71+int32(8), v65, v75, v66, v71+int32(12), int32(0), v68, int32(97))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_pfree(m, v75)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v109 = F_JsonbValueToJsonb(m, v105)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									m.G0 = v71 + int32(16)
									v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v114))) = v109
									m.G0 = v13 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
			v62 = F_pg_detoast_datum(m, v61)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				v64 = v62
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v68 = v13 + int32(28)
				v69 = m.G0
				v71 = v69 - int32(16)
				m.G0 = v71
				*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(0)
				v75 = F_palloc0(m, v66)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					if v77 != int32(16) {
					} else {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
						if v80 != int32(1) {
						} else {
							v84 = v13 + int32(36)
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
							v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							*(*int64)(unsafe.Add(mBase, uint32(v68))) = v86
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v88
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v84))) = v90
						}
					}
					v96 = F_JsonbIteratorInit(m, v64+int32(4))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v96
						v105 = F_setPath(m, v71+int32(8), v65, v75, v66, v71+int32(12), int32(0), v68, int32(97))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							F_pfree(m, v75)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return
							} else {
								v109 = F_JsonbValueToJsonb(m, v105)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									m.G0 = v71 + int32(16)
									v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v114))) = v109
									m.G0 = v13 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = v13 + int32(28)
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = int32(18)
			v29 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v23 + v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - v29
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
			if v39 == int32(1) {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v43 == int32(1) {
					v46 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+20)) = uint8(v46)
					v49 = int32(16)
				} else {
					v49 = int32(17)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v49
				v55 = F_JsonbValueToJsonb(m, v13+int32(8))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v58 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
					v64 = v55
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v68 = v13 + int32(28)
					v69 = m.G0
					v71 = v69 - int32(16)
					m.G0 = v71
					*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(0)
					v75 = F_palloc0(m, v66)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v77 != int32(16) {
						} else {
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
							if v80 != int32(1) {
							} else {
								v84 = v13 + int32(36)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								*(*int64)(unsafe.Add(mBase, uint32(v68))) = v86
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v88
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v84))) = v90
							}
						}
						v96 = F_JsonbIteratorInit(m, v64+int32(4))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v96
							v105 = F_setPath(m, v71+int32(8), v65, v75, v66, v71+int32(12), int32(0), v68, int32(97))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_pfree(m, v75)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v109 = F_JsonbValueToJsonb(m, v105)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										m.G0 = v71 + int32(16)
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = v109
										m.G0 = v13 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				v62 = F_pg_detoast_datum(m, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v64 = v62
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v68 = v13 + int32(28)
					v69 = m.G0
					v71 = v69 - int32(16)
					m.G0 = v71
					*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = int32(0)
					v75 = F_palloc0(m, v66)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						if v77 != int32(16) {
						} else {
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+12)))
							if v80 != int32(1) {
							} else {
								v84 = v13 + int32(36)
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
								v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								*(*int64)(unsafe.Add(mBase, uint32(v68))) = v86
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v68)+16)) = v88
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v84))) = v90
							}
						}
						v96 = F_JsonbIteratorInit(m, v64+int32(4))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v96
							v105 = F_setPath(m, v71+int32(8), v65, v75, v66, v71+int32(12), int32(0), v68, int32(97))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_pfree(m, v75)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return
								} else {
									v109 = F_JsonbValueToJsonb(m, v105)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										m.G0 = v71 + int32(16)
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = v109
										m.G0 = v13 + int32(48)
										return
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
func F_jsonb_to_recordset(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = int32(0)
	F_populate_recordset_worker(m, l0, int32(97429), v3, v3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
