package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v17 != 0 {
			v70 = v17
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v71 == int32(0) {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
				v79 = F_palloc(m, int32(12))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v79
					v82 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v82)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v79 + int32(2)
					if v75 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
						v98 = int64(0) - v75
						v101 = v82
						v103 = v79 + int32(12)
						v108 = v98
						for {
							v111 = v103 - int32(2)
							v113 = base.I64_div_u_s(v108, int64(10000))
							v116 = v113*int64(55536) + v108
							*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
							v119 = v101 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v108) {
								v101 = v119
								v103 = v111
								v108 = v113
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
						v123 = v119
						v128 = v101
					} else {
						v94 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v94
						if v75 == v94 {
							v123 = v82
							v128 = v2
						} else {
							v98 = v75
							v101 = v82
							v103 = v79 + int32(12)
							v108 = v98
							for {
								v111 = v103 - int32(2)
								v113 = base.I64_div_u_s(v108, int64(10000))
								v116 = v113*int64(55536) + v108
								*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
								v119 = v101 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v108) {
									v101 = v119
									v103 = v111
									v108 = v113
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
							v123 = v119
							v128 = v101
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v128
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v123
					v137 = F_make_result_opt_error(m, v12+int32(8), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v79)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_do_numeric_accum(m, v70, v137)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								m.G0 = v12 + int32(32)
								return v70
							}
						}
					}
				}
			} else {
				m.G0 = v12 + int32(32)
				return v70
			}
		} else {
			v20 = v12 + int32(8)
			v21 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v22 == v21 {
				v39 = int32(0)
				if v20 == v39 {
					v47 = v39
				} else {
					v42 = v39
					v43 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
				}
				v50 = v47
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				switch v25 - int32(429) {
				case 0:
					if v20 == int32(0) {
						v50 = int32(1)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
						v42 = v32
						v43 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
						v50 = v47
					}
				case 1:
					if v20 == int32(0) {
						v50 = int32(2)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+368))
						v42 = v37
						v43 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
						v50 = v47
					}
				default:
					v39 = int32(0)
					if v20 == v39 {
						v47 = v39
					} else {
						v42 = v39
						v43 = v21
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
					}
					v50 = v47
				}
			}
			if v50 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(58468), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476879), int32(4943), int32(336953))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v53 = int32(4442992)
				v54 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v56
				v59 = F_palloc0(m, int32(112))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v65
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v54
					v70 = v59
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v71 == int32(0) {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
						v79 = F_palloc(m, int32(12))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v79
							v82 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v82)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v79 + int32(2)
							if v75 < int64(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
								v98 = int64(0) - v75
								v101 = v82
								v103 = v79 + int32(12)
								v108 = v98
								for {
									v111 = v103 - int32(2)
									v113 = base.I64_div_u_s(v108, int64(10000))
									v116 = v113*int64(55536) + v108
									*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
									v119 = v101 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v108) {
										v101 = v119
										v103 = v111
										v108 = v113
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
								v123 = v119
								v128 = v101
							} else {
								v94 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v94
								if v75 == v94 {
									v123 = v82
									v128 = v2
								} else {
									v98 = v75
									v101 = v82
									v103 = v79 + int32(12)
									v108 = v98
									for {
										v111 = v103 - int32(2)
										v113 = base.I64_div_u_s(v108, int64(10000))
										v116 = v113*int64(55536) + v108
										*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
										v119 = v101 + int32(1)
										if base.Ui64(int64(9999)) < base.Ui64(v108) {
											v101 = v119
											v103 = v111
											v108 = v113
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
									v123 = v119
									v128 = v101
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v123
							v137 = F_make_result_opt_error(m, v12+int32(8), int32(0))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v79)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									F_do_numeric_accum(m, v70, v137)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										m.G0 = v12 + int32(32)
										return v70
									}
								}
							}
						}
					} else {
						m.G0 = v12 + int32(32)
						return v70
					}
				}
			}
		}
	} else {
		v20 = v12 + int32(8)
		v21 = int32(0)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v22 == v21 {
			v39 = int32(0)
			if v20 == v39 {
				v47 = v39
			} else {
				v42 = v39
				v43 = v21
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
				v47 = v43
			}
			v50 = v47
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			switch v25 - int32(429) {
			case 0:
				if v20 == int32(0) {
					v50 = int32(1)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					v42 = v32
					v43 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
					v50 = v47
				}
			case 1:
				if v20 == int32(0) {
					v50 = int32(2)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+368))
					v42 = v37
					v43 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
					v50 = v47
				}
			default:
				v39 = int32(0)
				if v20 == v39 {
					v47 = v39
				} else {
					v42 = v39
					v43 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
				}
				v50 = v47
			}
		}
		if v50 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v159 = m.ExcPending
			if v159 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(58468), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(476879), int32(4943), int32(336953))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v53 = int32(4442992)
			v54 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v56
			v59 = F_palloc0(m, int32(112))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v65
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v54
				v70 = v59
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v71 == int32(0) {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
					v79 = F_palloc(m, int32(12))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v79
						v82 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v82)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v79 + int32(2)
						if v75 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
							v98 = int64(0) - v75
							v101 = v82
							v103 = v79 + int32(12)
							v108 = v98
							for {
								v111 = v103 - int32(2)
								v113 = base.I64_div_u_s(v108, int64(10000))
								v116 = v113*int64(55536) + v108
								*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
								v119 = v101 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v108) {
									v101 = v119
									v103 = v111
									v108 = v113
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
							v123 = v119
							v128 = v101
						} else {
							v94 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v94
							if v75 == v94 {
								v123 = v82
								v128 = v2
							} else {
								v98 = v75
								v101 = v82
								v103 = v79 + int32(12)
								v108 = v98
								for {
									v111 = v103 - int32(2)
									v113 = base.I64_div_u_s(v108, int64(10000))
									v116 = v113*int64(55536) + v108
									*(*uint16)(unsafe.Add(mBase, uint32(v111))) = uint16(v116)
									v119 = v101 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v108) {
										v101 = v119
										v103 = v111
										v108 = v113
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v111
								v123 = v119
								v128 = v101
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v123
						v137 = F_make_result_opt_error(m, v12+int32(8), int32(0))
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v79)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_do_numeric_accum(m, v70, v137)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									m.G0 = v12 + int32(32)
									return v70
								}
							}
						}
					}
				} else {
					m.G0 = v12 + int32(32)
					return v70
				}
			}
		}
	}
}
func F_int8_bytea(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_cash_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int8_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
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
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v10 = F_PGLC_localeconv(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v14) {
			v17 = int32(2)
		} else {
			v17 = v14
		}
		v18 = base.I32_extend8_s(v17)
		if v18 <= int32(0) {
			v65 = int64(1)
		} else {
			if base.Ui32(v17) < base.Ui32(int32(8)) {
				v41 = int64(1)
			} else {
				v29 = int32(0)
				v30 = int64(1)
				for {
					v36 = v30 * int64(100000000)
					v38 = v29 + int32(8)
					if v38 != v18&int32(120) {
						v29 = v38
						v30 = v36
						continue
					} else {
						break
					}
					break
				}
				v41 = v36
			}
			if v17&int32(7) == int32(0) {
				v65 = v41
			} else {
				v53 = int32(0)
				v54 = v41
				for {
					v60 = v54 * int64(10)
					v62 = v53 + int32(1)
					if v62 != v18&int32(7) {
						v53 = v62
						v54 = v60
						continue
					} else {
						break
					}
					break
				}
				v65 = v60
			}
		}
		v72 = F_Int64GetDatum(m, v8)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			v74 = F_Int64GetDatum(m, v65)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = F_DirectFunctionCall2Coll(m, int32(1278), int32(0), v72, v74)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
					v79 = F_Int64GetDatum(m, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						return v79
					}
				}
			}
		}
	}
}
func F_int8_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v5 == int64(9223372036854775807) {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v8)
		return int32(0)
	} else {
		v12 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v12)
		v16 = F_Int64GetDatum(m, v5+int64(1))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			return v16
		}
	}
}
