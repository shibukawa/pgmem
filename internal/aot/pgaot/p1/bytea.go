package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_byteaGetByte(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v18 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v47 = v34
		} else {
			v35 = int32(1)
			if v18&v35 != 0 {
				v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		if base.B2i32(int32(0) <= v15)&base.B2i32(v15 < v47) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v47 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
					F_errmsg(m, int32(_a_F_byteaGetByte_0), v8)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaGetByte_1), int32(3325), int32(_a_F_byteaGetByte_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
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
			v71 = int32(1)
			if v18&v71 != 0 {
				v75 = v71
			} else {
				v75 = int32(4)
			}
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11+v75+v15))))
			m.G0 = v8 + int32(16)
			return v78
		}
	}
}
func F_byteaSetByte(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v17 = int32(base.Ui32(v15) >> (uint(int32(2)) % 32))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = int32(0)
		if base.B2i32(v19 <= v18)&base.B2i32(v18 < v17-int32(4)) == v19 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v17 - int32(5)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
					F_errmsg(m, int32(_a_F_byteaSetByte_0), v8)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_byteaSetByte_1), int32(3392), int32(_a_F_byteaSetByte_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*uint8)(unsafe.Add(mBase, uint32(v18+v11)+4)) = uint8(v47)
			m.G0 = v8 + int32(16)
			return v11
		}
	}
}
func F_bytea_int2(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if base.Ui32((v12-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_int2_1), int32(_a_F_bytea_int2_2), int32(_a_F_bytea_int2_3))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
				if v12 == int32(18) {
					v23 = int32(16)
				} else {
					v23 = int32(0)
				}
				v36 = v23
				if base.Ui32(int32(2)) < base.Ui32(v36) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_bytea_int2_1), int32(_a_F_bytea_int2_2), int32(_a_F_bytea_int2_3))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
					if v36 == int32(0) {
						return int32(0)
					} else {
						v43 = int32(1)
						if v9&v43 != 0 {
							v47 = v43
						} else {
							v47 = int32(4)
						}
						v48 = v5 + v47
						v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
						if v36 == int32(1) {
							return base.I32_extend16_s(v49)
						} else {
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
							return base.I32_extend16_s(v54 | v49<<(uint(int32(8))%32))
						}
					}
				}
			}
		} else {
			v24 = int32(1)
			if v9&v24 != 0 {
				v36 = int32(base.Ui32(v9)>>(uint(v24)%32)) - v24
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v36 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(2)) < base.Ui32(v36) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bytea_int2_0), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_int2_1), int32(_a_F_bytea_int2_2), int32(_a_F_bytea_int2_3))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
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
				if v36 == int32(0) {
					return int32(0)
				} else {
					v43 = int32(1)
					if v9&v43 != 0 {
						v47 = v43
					} else {
						v47 = int32(4)
					}
					v48 = v5 + v47
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
					if v36 == int32(1) {
						return base.I32_extend16_s(v49)
					} else {
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
						return base.I32_extend16_s(v54 | v49<<(uint(int32(8))%32))
					}
				}
			}
		}
	}
}
func F_bytea_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v13 == int32(1) {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v47 = int32(4)
				v50 = v47 & int32(3)
				v51 = int32(1)
				if v13&v51 != 0 {
					v55 = v51
				} else {
					v55 = int32(4)
				}
				v56 = v9 + v55
				v57 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v47) {
					v63 = v57
					v64 = v57
					for {
						v70 = int32(4)
						v71 = v63 + v70
						v73 = v64 + v70
						if v73 != v47&int32(4) {
							v63 = v71
							v64 = v73
							continue
						} else {
							break
						}
						break
					}
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
					v77 = int32(16711935)
					v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
					if v50 == int32(0) {
						v113 = v85
					} else {
						v88 = v71
						v89 = v85
						v95 = v88
						v96 = v89
						v100 = int32(0)
						for {
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
							v106 = v103 | v96<<(uint(int32(8))%32)
							v107 = int32(1)
							v110 = v100 + v107
							if v110 != v50 {
								v95 = v95 + v107
								v96 = v106
								v100 = v110
								continue
							} else {
								break
							}
							break
						}
						v113 = v106
					}
				} else {
					v88 = v57
					v89 = v57
					v95 = v88
					v96 = v89
					v100 = int32(0)
					for {
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
						v106 = v103 | v96<<(uint(int32(8))%32)
						v107 = int32(1)
						v110 = v100 + v107
						if v110 != v50 {
							v95 = v95 + v107
							v96 = v106
							v100 = v110
							continue
						} else {
							break
						}
						break
					}
					v113 = v106
				}
				return v113
			} else {
				if v17 == int32(18) {
					v28 = int32(16)
				} else {
					v28 = int32(0)
				}
				v42 = v28
				if base.Ui32(int32(4)) < base.Ui32(v42) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int4_0), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_bytea_int4_1), int32(_a_F_bytea_int4_2), int32(_a_F_bytea_int4_3))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
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
					if v42 != 0 {
						v47 = v42
						v50 = v47 & int32(3)
						v51 = int32(1)
						if v13&v51 != 0 {
							v55 = v51
						} else {
							v55 = int32(4)
						}
						v56 = v9 + v55
						v57 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v47) {
							v63 = v57
							v64 = v57
							for {
								v70 = int32(4)
								v71 = v63 + v70
								v73 = v64 + v70
								if v73 != v47&int32(4) {
									v63 = v71
									v64 = v73
									continue
								} else {
									break
								}
								break
							}
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
							v77 = int32(16711935)
							v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
							if v50 == int32(0) {
								v113 = v85
							} else {
								v88 = v71
								v89 = v85
								v95 = v88
								v96 = v89
								v100 = int32(0)
								for {
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
									v106 = v103 | v96<<(uint(int32(8))%32)
									v107 = int32(1)
									v110 = v100 + v107
									if v110 != v50 {
										v95 = v95 + v107
										v96 = v106
										v100 = v110
										continue
									} else {
										break
									}
									break
								}
								v113 = v106
							}
						} else {
							v88 = v57
							v89 = v57
							v95 = v88
							v96 = v89
							v100 = int32(0)
							for {
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
								v106 = v103 | v96<<(uint(int32(8))%32)
								v107 = int32(1)
								v110 = v100 + v107
								if v110 != v50 {
									v95 = v95 + v107
									v96 = v106
									v100 = v110
									continue
								} else {
									break
								}
								break
							}
							v113 = v106
						}
						return v113
					} else {
						return int32(0)
					}
				}
			}
		} else {
			v29 = int32(1)
			if v13&v29 != 0 {
				v42 = int32(base.Ui32(v13)>>(uint(v29)%32)) - v29
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(4)) < base.Ui32(v42) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bytea_int4_0), int32(0))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_int4_1), int32(_a_F_bytea_int4_2), int32(_a_F_bytea_int4_3))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
				if v42 != 0 {
					v47 = v42
					v50 = v47 & int32(3)
					v51 = int32(1)
					if v13&v51 != 0 {
						v55 = v51
					} else {
						v55 = int32(4)
					}
					v56 = v9 + v55
					v57 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v47) {
						v63 = v57
						v64 = v57
						for {
							v70 = int32(4)
							v71 = v63 + v70
							v73 = v64 + v70
							if v73 != v47&int32(4) {
								v63 = v71
								v64 = v73
								continue
							} else {
								break
							}
							break
						}
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v63+v56)))
						v77 = int32(16711935)
						v85 = base.I32_rotr(v76&v77, int32(8)) | base.I32_rotr(v76, int32(24))&v77
						if v50 == int32(0) {
							v113 = v85
						} else {
							v88 = v71
							v89 = v85
							v95 = v88
							v96 = v89
							v100 = int32(0)
							for {
								v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
								v106 = v103 | v96<<(uint(int32(8))%32)
								v107 = int32(1)
								v110 = v100 + v107
								if v110 != v50 {
									v95 = v95 + v107
									v96 = v106
									v100 = v110
									continue
								} else {
									break
								}
								break
							}
							v113 = v106
						}
					} else {
						v88 = v57
						v89 = v57
						v95 = v88
						v96 = v89
						v100 = int32(0)
						for {
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95+v56))))
							v106 = v103 | v96<<(uint(int32(8))%32)
							v107 = int32(1)
							v110 = v100 + v107
							if v110 != v50 {
								v95 = v95 + v107
								v96 = v106
								v100 = v110
								continue
							} else {
								break
							}
							break
						}
						v113 = v106
					}
					return v113
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_bytea_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v84 int64
	_ = v84
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v2 = int32(0)
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v14 == int32(1) {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v50 = int32(4)
				v53 = v50 & int32(3)
				v54 = int32(1)
				if v14&v54 != 0 {
					v58 = v54
				} else {
					v58 = int32(4)
				}
				v59 = v10 + v58
				v60 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v50) {
					v65 = v60
					v71 = v2
					v72 = v8
					for {
						v73 = int64(16)
						v75 = v65 + v59
						v76 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						v77 = int64(8)
						v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
						v84 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
						v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
						v89 = (v72<<(uint(v73)%64)|v76<<(uint(v77)%64)|v80)<<(uint(v73)%64) | v84<<(uint(v77)%64) | v88
						v90 = int32(4)
						v91 = v65 + v90
						v93 = v71 + v90
						if v93 != v50&int32(12) {
							v65 = v91
							v71 = v93
							v72 = v89
							continue
						} else {
							break
						}
						break
					}
					if v53 == int32(0) {
						v130 = v89
					} else {
						v97 = v91
						v104 = v89
						v105 = v97
						v109 = v2
						v112 = v104
						for {
							v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
							v117 = v114 | v112<<(uint(int64(8))%64)
							v118 = int32(1)
							v121 = v109 + v118
							if v121 != v53 {
								v105 = v105 + v118
								v109 = v121
								v112 = v117
								continue
							} else {
								break
							}
							break
						}
						v130 = v117
					}
				} else {
					v97 = v60
					v104 = v8
					v105 = v97
					v109 = v2
					v112 = v104
					for {
						v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
						v117 = v114 | v112<<(uint(int64(8))%64)
						v118 = int32(1)
						v121 = v109 + v118
						if v121 != v53 {
							v105 = v105 + v118
							v109 = v121
							v112 = v117
							continue
						} else {
							break
						}
						break
					}
					v130 = v117
				}
				v131 = F_Int64GetDatum(m, v130)
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					return v131
				}
			} else {
				if v18 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				v43 = v29
				if base.Ui32(int32(8)) < base.Ui32(v43) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_bytea_int8_0), int32(0))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_bytea_int8_1), int32(_a_F_bytea_int8_2), int32(_a_F_bytea_int8_3))
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
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
					if v43 != 0 {
						v50 = v43
						v53 = v50 & int32(3)
						v54 = int32(1)
						if v14&v54 != 0 {
							v58 = v54
						} else {
							v58 = int32(4)
						}
						v59 = v10 + v58
						v60 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v50) {
							v65 = v60
							v71 = v2
							v72 = v8
							for {
								v73 = int64(16)
								v75 = v65 + v59
								v76 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
								v77 = int64(8)
								v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
								v84 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
								v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
								v89 = (v72<<(uint(v73)%64)|v76<<(uint(v77)%64)|v80)<<(uint(v73)%64) | v84<<(uint(v77)%64) | v88
								v90 = int32(4)
								v91 = v65 + v90
								v93 = v71 + v90
								if v93 != v50&int32(12) {
									v65 = v91
									v71 = v93
									v72 = v89
									continue
								} else {
									break
								}
								break
							}
							if v53 == int32(0) {
								v130 = v89
							} else {
								v97 = v91
								v104 = v89
								v105 = v97
								v109 = v2
								v112 = v104
								for {
									v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
									v117 = v114 | v112<<(uint(int64(8))%64)
									v118 = int32(1)
									v121 = v109 + v118
									if v121 != v53 {
										v105 = v105 + v118
										v109 = v121
										v112 = v117
										continue
									} else {
										break
									}
									break
								}
								v130 = v117
							}
						} else {
							v97 = v60
							v104 = v8
							v105 = v97
							v109 = v2
							v112 = v104
							for {
								v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
								v117 = v114 | v112<<(uint(int64(8))%64)
								v118 = int32(1)
								v121 = v109 + v118
								if v121 != v53 {
									v105 = v105 + v118
									v109 = v121
									v112 = v117
									continue
								} else {
									break
								}
								break
							}
							v130 = v117
						}
						v131 = F_Int64GetDatum(m, v130)
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							return v131
						}
					} else {
						v47 = F_Int64GetDatum(m, int64(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							return v47
						}
					}
				}
			}
		} else {
			v30 = int32(1)
			if v14&v30 != 0 {
				v43 = int32(base.Ui32(v14)>>(uint(v30)%32)) - v30
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v43 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
			}
			if base.Ui32(int32(8)) < base.Ui32(v43) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bytea_int8_0), int32(0))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bytea_int8_1), int32(_a_F_bytea_int8_2), int32(_a_F_bytea_int8_3))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
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
				if v43 != 0 {
					v50 = v43
					v53 = v50 & int32(3)
					v54 = int32(1)
					if v14&v54 != 0 {
						v58 = v54
					} else {
						v58 = int32(4)
					}
					v59 = v10 + v58
					v60 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v50) {
						v65 = v60
						v71 = v2
						v72 = v8
						for {
							v73 = int64(16)
							v75 = v65 + v59
							v76 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v77 = int64(8)
							v80 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
							v84 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+2)))
							v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+3)))
							v89 = (v72<<(uint(v73)%64)|v76<<(uint(v77)%64)|v80)<<(uint(v73)%64) | v84<<(uint(v77)%64) | v88
							v90 = int32(4)
							v91 = v65 + v90
							v93 = v71 + v90
							if v93 != v50&int32(12) {
								v65 = v91
								v71 = v93
								v72 = v89
								continue
							} else {
								break
							}
							break
						}
						if v53 == int32(0) {
							v130 = v89
						} else {
							v97 = v91
							v104 = v89
							v105 = v97
							v109 = v2
							v112 = v104
							for {
								v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
								v117 = v114 | v112<<(uint(int64(8))%64)
								v118 = int32(1)
								v121 = v109 + v118
								if v121 != v53 {
									v105 = v105 + v118
									v109 = v121
									v112 = v117
									continue
								} else {
									break
								}
								break
							}
							v130 = v117
						}
					} else {
						v97 = v60
						v104 = v8
						v105 = v97
						v109 = v2
						v112 = v104
						for {
							v114 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v105+v59))))
							v117 = v114 | v112<<(uint(int64(8))%64)
							v118 = int32(1)
							v121 = v109 + v118
							if v121 != v53 {
								v105 = v105 + v118
								v109 = v121
								v112 = v117
								continue
							} else {
								break
							}
							break
						}
						v130 = v117
					}
					v131 = F_Int64GetDatum(m, v130)
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return int32(0)
					} else {
						return v131
					}
				} else {
					v47 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						return v47
					}
				}
			}
		}
	}
}
func F_bytea_larger(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v15 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v45 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v21 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v32 = int32(1)
	if v15&v32 != 0 {
		v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v24 = int32(16)
	goto L10
L9:
	;
	v24 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = int32(4)
	goto L13
L12:
	;
	v31 = v24
	goto L13
L13:
	;
	v44 = v31
	goto L4
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v74 < v44 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v51 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v62 = int32(1)
	if v45&v62 != 0 {
		v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v54 = int32(16)
	goto L21
L20:
	;
	v54 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(4)
	goto L24
L23:
	;
	v61 = v54
	goto L24
L24:
	;
	v74 = v61
	goto L15
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v76 = v8
	goto L28
L27:
	;
	v76 = v13
	goto L28
L28:
	;
	v77 = int32(1)
	if v15&v77 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v81 = v77
	goto L31
L30:
	;
	v81 = int32(4)
	goto L31
L31:
	;
	v82 = v8 + v81
	v83 = int32(1)
	if v45&v83 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = v83
	goto L34
L33:
	;
	v87 = int32(4)
	goto L34
L34:
	;
	v88 = v13 + v87
	if v44 < v74 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v44
	goto L37
L36:
	;
	v90 = v74
	goto L37
L37:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v152 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v152 = int32(0)
	goto L38
L40:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L50
L41:
	;
	if (v82|v88)&int32(3) != 0 {
		v121 = v82
		v122 = v88
		v123 = v90
		goto L40
	} else {
		goto L44
	}
L42:
	;
	v114 = v82
	v115 = v88
	v116 = v90
	goto L43
L43:
	;
	if v116 == int32(0) {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v98 = v82
	v99 = v88
	v100 = v90
	goto L45
L45:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L40
	} else {
		goto L47
	}
L46:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L43
L47:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L40
L50:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v152 = v131 - v132
	goto L38
L52:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	goto L39
L56:
	;
	v153 = v13
	goto L58
L57:
	;
	v153 = v76
	goto L58
L58:
	;
	if int32(0) < v152 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v156 = v8
	goto L61
L60:
	;
	v156 = v153
	goto L61
L61:
	;
	return v156
}
func F_bytea_substr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = F_bytea_substring(m, v2, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_bytea_substring(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	if int32(0) <= l2 {
		v6 = int32(1)
		if l1 <= v6 {
			v9 = v6
		} else {
			v9 = l1
		}
		v14 = l1 + l2
		if base.B2i32(l2 < int32(0))^base.B2i32(v14 < l1) != 0 {
			v31 = int32(-1)
			v32 = F_detoast_attr_slice(m, l0, v9-int32(1), v31)
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return v32
			}
		} else {
			if v14 <= int32(0) {
				v23 = F_DirectFunctionCall1Coll(m, int32(579), int32(0), int32(_a_F_bytea_substring_0))
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_pg_detoast_datum_packed(m, v23)
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			} else {
				v31 = v14 - v9
				v32 = F_detoast_attr_slice(m, l0, v9-int32(1), v31)
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					return v32
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(17039490))
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_bytea_substring_1), int32(0))
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_bytea_substring_2), int32(3158), int32(_a_F_bytea_substring_3))
					v50 = m.ExcPending
					if v50 != 0 {
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
