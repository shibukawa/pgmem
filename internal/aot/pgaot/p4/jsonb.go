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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
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
						v79 = v7 + int32(16)
						F_initStringInfo(m, v79)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							F_enlargeStringInfo(m, v79, int32(4))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
								v87 = v85 + int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								v91 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
								F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
									v147 = v98
									m.G0 = v7 + int32(32)
									return v147
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
								v79 = v7 + int32(16)
								F_initStringInfo(m, v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v79, int32(4))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										v87 = v85 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v91 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
										F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
											v147 = v98
											m.G0 = v7 + int32(32)
											return v147
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
							v79 = v7 + int32(16)
							F_initStringInfo(m, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v79, int32(4))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									v87 = v85 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
									F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
										v147 = v98
										m.G0 = v7 + int32(32)
										return v147
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
								F_errmsg_internal(m, int32(_a_F_JsonbValueToJsonb_0), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_JsonbValueToJsonb_1), int32(720), int32(_a_F_JsonbValueToJsonb_2))
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
			v104 = v7 + int32(16)
			F_initStringInfo(m, v104)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_enlargeStringInfo(m, v104, int32(4))
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
					v112 = v110 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v112
					v114 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
					v116 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v114+v112))) = uint8(v116)
					F_convertJsonbValue(m, v104, v7+int32(12), l0, v116)
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int32(0)
					} else {
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v123))) = v124 << (uint(int32(2)) % 32)
						v147 = v123
						m.G0 = v7 + int32(32)
						return v147
					}
				}
			}
		default:
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v131 = F_palloc(m, v128+int32(4))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return int32(0)
			} else {
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v131))) = v133<<(uint(int32(2))%32) + int32(16)
				v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v139 == int32(0) {
					v147 = v131
				} else {
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					base.MemoryCopy(m, v131+int32(4), v144, v139)
					v147 = v131
				}
				m.G0 = v7 + int32(32)
				return v147
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
							v79 = v7 + int32(16)
							F_initStringInfo(m, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_enlargeStringInfo(m, v79, int32(4))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									v87 = v85 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
									v91 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
									F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
										v147 = v98
										m.G0 = v7 + int32(32)
										return v147
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
									v79 = v7 + int32(16)
									F_initStringInfo(m, v79)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										F_enlargeStringInfo(m, v79, int32(4))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											v87 = v85 + int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v91 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
											F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
												v147 = v98
												m.G0 = v7 + int32(32)
												return v147
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
								v79 = v7 + int32(16)
								F_initStringInfo(m, v79)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_enlargeStringInfo(m, v79, int32(4))
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int32(0)
									} else {
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
										v87 = v85 + int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v87
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v91 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v89+v87))) = uint8(v91)
										F_convertJsonbValue(m, v79, v7+int32(12), v39, v91)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 << (uint(int32(2)) % 32)
											v147 = v98
											m.G0 = v7 + int32(32)
											return v147
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
									F_errmsg_internal(m, int32(_a_F_JsonbValueToJsonb_0), int32(0))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_JsonbValueToJsonb_1), int32(720), int32(_a_F_JsonbValueToJsonb_2))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13933(m, l0, int32(5), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13933(m, l0, int32(7), int32(6))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
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
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
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
	v320 = m.ExcPending
	if v320 != 0 {
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
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L76
	}
L8:
	;
	if v32&int32(268435455) == int32(0) {
		v285 = v20
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v17 + int32(48)
	return v285
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
		v285 = v20
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
	v281 = F_JsonbValueToJsonb(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L75
	}
L15:
	;
	if v61 == int32(0) {
		v280 = v2
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v65 = v61
	v78 = v2
	goto L17
L17:
	;
	v80 = base.B2i32(v65 != int32(1))
	if v80&base.B2i32(v65 != int32(3)) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v280 = v78
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
	v239 = v17 + int32(8)
	if base.Ui32(v65) < base.Ui32(int32(4)) {
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
	v220 = v97 + int32(1)
	if v220 != v87 {
		v97 = v220
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
	if v146 != v92 {
		goto L26
	} else {
		goto L39
	}
L29:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v125 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v136 = int32(1)
	if v119 != 0 {
		v146 = int32(base.Ui32(v117)>>(uint(v136)%32)) - v136
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v128 = int32(16)
	goto L34
L33:
	;
	v128 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v125-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = int32(4)
	goto L37
L36:
	;
	v135 = v128
	goto L37
L37:
	;
	v146 = v135
	goto L28
L38:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v146 = int32(base.Ui32(v140)>>(uint(int32(2))%32)) - int32(4)
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
	v150 = v116
	goto L42
L41:
	;
	v150 = v114 + int32(4)
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
	if v212 == int32(0) {
		goto L19
	} else {
		goto L61
	}
L44:
	;
	v212 = int32(0)
	goto L43
L45:
	;
	v186 = v181
	v187 = v182
	v188 = v183
	goto L55
L46:
	;
	if (v150|v91)&int32(3) != 0 {
		v181 = v150
		v182 = v91
		v183 = v92
		goto L45
	} else {
		goto L49
	}
L47:
	;
	v174 = v150
	v175 = v91
	v176 = v92
	goto L48
L48:
	;
	if v176 == int32(0) {
		goto L44
	} else {
		goto L54
	}
L49:
	;
	v158 = v150
	v159 = v91
	v160 = v92
	goto L50
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	if v163 != v164 {
		v181 = v158
		v182 = v159
		v183 = v160
		goto L45
	} else {
		goto L52
	}
L51:
	;
	v174 = v169
	v175 = v167
	v176 = v171
	goto L48
L52:
	;
	v166 = int32(4)
	v167 = v159 + v166
	v169 = v158 + v166
	v171 = v160 - v166
	if base.Ui32(int32(3)) < base.Ui32(v171) {
		v158 = v169
		v159 = v167
		v160 = v171
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v181 = v174
	v182 = v175
	v183 = v176
	goto L45
L55:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v191 == v192 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v212 = v191 - v192
	goto L43
L57:
	;
	v194 = int32(1)
	v199 = v188 - v194
	if v199 != 0 {
		v186 = v186 + v194
		v187 = v187 + v194
		v188 = v199
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
	v243 = v239
	goto L65
L64:
	;
	v243 = int32(0)
	goto L65
L65:
	;
	v244 = F_pushJsonbValue(m, v17+int32(32), v65, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v249 = F_JsonbIteratorNext(m, v17+int32(28), v239, int32(1))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v249 != 0 {
		v65 = v249
		v78 = v244
		goto L17
	} else {
		goto L68
	}
L68:
	;
	v280 = v244
	goto L14
L69:
	;
	v258 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v265 = F_JsonbIteratorNext(m, v17+int32(28), v17+int32(8), int32(1))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	if v265 != 0 {
		v65 = v265
		goto L17
	} else {
		goto L74
	}
L74:
	;
	goto L18
L75:
	;
	v285 = v281
	goto L9
L76:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_array_0), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_array_1), int32(_a_F_jsonb_delete_array_2), int32(_a_F_jsonb_delete_array_3))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
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
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_array_4), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_array_1), int32(_a_F_jsonb_delete_array_5), int32(_a_F_jsonb_delete_array_3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
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
	var v37 int32
	_ = v37
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
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(1387)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(1388)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = int32(1389)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1390)
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
	v37 = int32(0)
	goto L6
L6:
	;
	v40 = v37 << (uint(int32(2)) % 32)
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
	v50 = v37 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v50 < v51 {
		v37 = v50
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
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
	return v238
L2:
	;
	if l4 == int32(0) {
		v238 = l0
		goto L1
	} else {
		goto L72
	}
L3:
	;
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v225)
	v238 = int32(0)
	goto L1
L4:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v222)
	v238 = int32(0)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
	v54 = int32(0)
	v60 = v49
	v61 = v20
	v62 = v26
	goto L16
L8:
	;
	v43 = int32(0)
	v46 = base.B2i32(l2 <= v43)
	if base.B2i32(v42 == v43)&v46 != 0 {
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
		v49 = v33
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
	if l2 <= v43 {
		v191 = v42
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v49 = v41
	goto L7
L16:
	;
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v158 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1+v54<<(uint(int32(2))%32))))
	v70 = F_pg_detoast_datum_packed(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v60 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L22:
	;
	v72 = int32(1)
	v73 = v70 + v72
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v78 = v76 & v72
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = v73
	goto L25
L24:
	;
	v79 = v70 + int32(4)
	goto L25
L25:
	;
	if v76 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v85 == int32(18) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v78 != 0 {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v88 = int32(16)
	goto L31
L30:
	;
	v88 = int32(0)
	goto L31
L31:
	;
	if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = int32(4)
	goto L34
L33:
	;
	v95 = v88
	goto L34
L34:
	;
	v97 = F_getKeyJsonValueFromContainer(m, v61, v79, v95, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v158 = v97
	goto L18
L36:
	;
	v99 = int32(1)
	v104 = F_getKeyJsonValueFromContainer(m, v61, v79, int32(base.Ui32(v76)>>(uint(v99)%32))-v99, int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L12
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v112 = F_getKeyJsonValueFromContainer(m, v61, v79, int32(base.Ui32(v106)>>(uint(int32(2))%32))-int32(4), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	v158 = v104
	goto L18
L40:
	;
	v158 = v112
	goto L18
L41:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1+v54<<(uint(int32(2))%32))))
	v120 = F_text_to_cstring(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_get_element[0])) = int32(0)
	v128 = F_strtol(m, v120, v15+int32(12), int32(10))
	mBase = m.M
	goto L43
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v120 == v129 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	if v131 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_get_element[0]))
	if v133 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v128 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v136&int32(1073741824) == int32(0) {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v150 = v128
	goto L49
L49:
	;
	v151 = F_getIthJsonbValueFromContainer(m, v61, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L53
	}
L50:
	;
	if v128 == int32(-2147483648) {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v144 = v136 & int32(268435455)
	if base.Ui32(v144) < base.Ui32(int32(0)-v128) {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v150 = v128 + v144
	goto L49
L53:
	;
	v158 = v151
	goto L18
L54:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v161)
	v238 = int32(0)
	goto L1
L55:
	;
	goto L56
L56:
	;
	if v54 == l2-int32(1) {
		v191 = v158
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v167 == int32(18) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v181 = int32(base.Ui32(v171&int32(1073741824)) >> (uint(int32(30)) % 32))
	v182 = v170
	v183 = int32(base.Ui32(v171&int32(536870912)) >> (uint(int32(29)) % 32))
	goto L60
L59:
	;
	v181 = v165
	v182 = v61
	v183 = v165
	goto L60
L60:
	;
	v54 = v54 + int32(1)
	v60 = v181
	v61 = v182
	v62 = v183
	goto L16
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v198 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v206 = F_JsonbValueToJsonb(m, v191)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L68
	}
L64:
	;
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v201)
	v238 = int32(0)
	goto L1
L65:
	;
	goto L66
L66:
	;
	v204 = F_JsonbValueAsText(m, v191)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	v238 = v204
	goto L1
L68:
	;
	v238 = v206
	goto L1
L69:
	;
	F_errmsg_internal(m, int32(_a_F_jsonb_get_element_0), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_jsonb_get_element_1), int32(1614), int32(_a_F_jsonb_get_element_2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
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
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = F_JsonbToCString(m, int32(0), v20, int32(base.Ui32(v231)>>(uint(int32(2))%32)))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v236 = F_cstring_to_text(m, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v238 = v236
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13934(m, l0, int32(_a_F_jsonb_int4_0), int32(1318))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v17 <= int32(0))
							}
						} else {
							return base.B2i32(v17 <= int32(0))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v17 <= int32(0))
						}
					} else {
						return base.B2i32(v17 <= int32(0))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)))
	if v16 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(0)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
		if v38 == int32(1) {
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			if v42 == int32(1) {
				v45 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v45)
				v48 = int32(16)
			} else {
				v48 = int32(17)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v48
			v54 = F_JsonbValueToJsonb(m, v12+int32(8))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v57 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
				v63 = v54
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v67 = v12 + int32(28)
				v68 = m.G0
				v70 = v68 - int32(16)
				m.G0 = v70
				*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(0)
				v74 = F_palloc0(m, v65)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					if v76 != int32(16) {
					} else {
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+12)))
						if v79 != int32(1) {
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
							*(*int64)(unsafe.Add(mBase, uint32(v67))) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v85
							v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v87
						}
					}
					v92 = F_JsonbIteratorInit(m, v63+int32(4))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v92
						v101 = F_setPath(m, v70+int32(8), v64, v74, v65, v70+int32(12), int32(0), v67, int32(97))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_pfree(m, v74)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								v105 = F_JsonbValueToJsonb(m, v101)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									m.G0 = v70 + int32(16)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v110))) = v105
									m.G0 = v12 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v61 = F_pg_detoast_datum(m, v60)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				v63 = v61
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v67 = v12 + int32(28)
				v68 = m.G0
				v70 = v68 - int32(16)
				m.G0 = v70
				*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(0)
				v74 = F_palloc0(m, v65)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
					if v76 != int32(16) {
					} else {
						v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+12)))
						if v79 != int32(1) {
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
							v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
							*(*int64)(unsafe.Add(mBase, uint32(v67))) = v83
							v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v85
							v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v87
						}
					}
					v92 = F_JsonbIteratorInit(m, v63+int32(4))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v92
						v101 = F_setPath(m, v70+int32(8), v64, v74, v65, v70+int32(12), int32(0), v67, int32(97))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_pfree(m, v74)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								v105 = F_JsonbValueToJsonb(m, v101)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									m.G0 = v70 + int32(16)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v110))) = v105
									m.G0 = v12 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v25 = v12 + int32(28)
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(18)
			v28 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v22 + v28
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - v28
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
			if v38 == int32(1) {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v42 == int32(1) {
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v45)
					v48 = int32(16)
				} else {
					v48 = int32(17)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v48
				v54 = F_JsonbValueToJsonb(m, v12+int32(8))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v57)
					v63 = v54
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v67 = v12 + int32(28)
					v68 = m.G0
					v70 = v68 - int32(16)
					m.G0 = v70
					*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(0)
					v74 = F_palloc0(m, v65)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
						if v76 != int32(16) {
						} else {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+12)))
							if v79 != int32(1) {
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
								*(*int64)(unsafe.Add(mBase, uint32(v67))) = v83
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v85
								v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v87
							}
						}
						v92 = F_JsonbIteratorInit(m, v63+int32(4))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v92
							v101 = F_setPath(m, v70+int32(8), v64, v74, v65, v70+int32(12), int32(0), v67, int32(97))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								F_pfree(m, v74)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return
								} else {
									v105 = F_JsonbValueToJsonb(m, v101)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										m.G0 = v70 + int32(16)
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v110))) = v105
										m.G0 = v12 + int32(48)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				v61 = F_pg_detoast_datum(m, v60)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v63 = v61
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v67 = v12 + int32(28)
					v68 = m.G0
					v70 = v68 - int32(16)
					m.G0 = v70
					*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = int32(0)
					v74 = F_palloc0(m, v65)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
						if v76 != int32(16) {
						} else {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+12)))
							if v79 != int32(1) {
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
								v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
								*(*int64)(unsafe.Add(mBase, uint32(v67))) = v83
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v85
								v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v87
							}
						}
						v92 = F_JsonbIteratorInit(m, v63+int32(4))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v92
							v101 = F_setPath(m, v70+int32(8), v64, v74, v65, v70+int32(12), int32(0), v67, int32(97))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								F_pfree(m, v74)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return
								} else {
									v105 = F_JsonbValueToJsonb(m, v101)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										m.G0 = v70 + int32(16)
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v110))) = v105
										m.G0 = v12 + int32(48)
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
	F_populate_recordset_worker(m, l0, int32(_a_F_jsonb_to_recordset_0), v3, v3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
