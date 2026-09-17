package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckRelationTableSpaceMove(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	if l1 == v12 {
		v34 = v3
		m.G0 = v9 + int32(16)
		return v34
	} else {
		if v12 == int32(0) {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_CheckRelationTableSpaceMove[0]))
			if l1 == v17 {
				v34 = v3
				m.G0 = v9 + int32(16)
				return v34
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
				switch v19 - int32(83) {
				case 0, 22, 26, 31, 33:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
					if v22 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_0), v9)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3714), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						if l1 == int32(1664) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3720), int32(_a_F_CheckRelationTableSpaceMove_2))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
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
							v27 = int32(1)
							v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
							if v28 != int32(116) {
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							} else {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
								if v31 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3729), int32(_a_F_CheckRelationTableSpaceMove_2))
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
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
									v34 = v27
									m.G0 = v9 + int32(16)
									return v34
								}
							}
						}
					}
				default:
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3720), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3729), int32(_a_F_CheckRelationTableSpaceMove_2))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
			switch v19 - int32(83) {
			case 0, 22, 26, 31, 33:
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+88))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v48 + int32(4)
							F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_0), v9)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3714), int32(_a_F_CheckRelationTableSpaceMove_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
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
					if l1 == int32(1664) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3720), int32(_a_F_CheckRelationTableSpaceMove_2))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
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
						v27 = int32(1)
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
						if v28 != int32(116) {
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
							if v31 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3729), int32(_a_F_CheckRelationTableSpaceMove_2))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								v34 = v27
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			default:
				if l1 == int32(1664) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_3), int32(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3720), int32(_a_F_CheckRelationTableSpaceMove_2))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
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
					v27 = int32(1)
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+118)))
					if v28 != int32(116) {
						v34 = v27
						m.G0 = v9 + int32(16)
						return v34
					} else {
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckRelationTableSpaceMove_4), int32(0))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_CheckRelationTableSpaceMove_1), int32(3729), int32(_a_F_CheckRelationTableSpaceMove_2))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
							v34 = v27
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_GetRelationPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(224)
	m.G0 = v10
	switch l2 - int32(1663) {
	case 0:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+180)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = l1
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+184)) = v40
				v45 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_0), v10+int32(176))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+164)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = l1
				v52 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_1), v10+int32(160))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+216)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+212)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = l1
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+220)) = v61
				v66 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_2), v10+int32(208))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+192)) = l1
				v74 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_3), v10+int32(192))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	case 1:
		if l5 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+144)) = l3
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+148)) = v19
			v24 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_4), v10+int32(144))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+128)) = l3
			v30 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_5), v10+int32(128))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v10 + int32(224)
				return
			}
		}
	default:
		if l4 == int32(-1) {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l3
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v83
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_GetRelationPath_7)
				v94 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_8), v10+int32(32))
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_GetRelationPath_7)
				v104 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_9), v10)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		} else {
			if l5 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+116)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l4
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l5<<(uint(int32(2))%32))+uint32(_c_F_GetRelationPath[0])))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v112
				*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+104)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = int32(_a_F_GetRelationPath_7)
				v123 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_10), v10+int32(96))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+76)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = int32(_a_F_GetRelationPath_6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = int32(_a_F_GetRelationPath_7)
				v136 = F_pg_sprintf(m, l0, int32(_a_F_GetRelationPath_11), v10-int32(-64))
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					m.G0 = v10 + int32(224)
					return
				}
			}
		}
	}
}
func F_LockRelationOid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_LockRelationOid_0) {
			v36 = l0 - int32(_a_F_LockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_LockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_LockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_LockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_LockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_LockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v89
	v93 = int32(0)
	v98 = F_LockAcquireExtended(m, v7+int32(16), l1, v93, v93, v7+int32(12), v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		return
	} else {
		if v98 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				v105 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v104)+53)) = uint8(v105)
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	}
}
func F_RelationBuildDesc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int64
	_ = v111
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v936 int32
	_ = v936
	var v942 int64
	_ = v942
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1161 int32
	_ = v1161
	v20 = m.G0
	v22 = v20 - int32(320)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[2]))
	if v29 <= v27 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = F_repalloc(m, v25, v29<<(uint(int32(4))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v45 = v27
	v46 = v25
	goto L3
L3:
	;
	v48 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v45 + v48
	v52 = v45 << (uint(int32(3)) % 32)
	v53 = v46 + v52
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0
	v59 = F_ScanPgRelation(m, l0, v48, v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[2])) = v29 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0])) = v33
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	v45 = v44
	v46 = v33
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L259
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L4
	} else {
		goto L256
	}
L8:
	;
	m.G0 = v22 + int32(320)
	return v1104
L9:
	;
	v1047 = int32(_a_F_RelationBuildDesc_0)
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v1049 - int32(1)
	if l1 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L10:
	;
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v77 = v59
	goto L14
L12:
	;
	goto L13
L13:
	;
	v1041 = int32(_a_F_RelationBuildDesc_0)
	v1043 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[1])) = v1043 - int32(1)
	v1104 = int32(0)
	goto L8
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v84 = v82 + v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = int32(_a_F_RelationBuildDesc_1)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3]))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v90
	v93 = F_palloc0(m, int32(276))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = int32(0)
	v98 = F_palloc(m, int32(144))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	base.MemoryCopy(m, v98, v84, int32(144))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+48)) = v98
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98)+120)))
	v104 = F_CreateTemplateTupleDesc(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+52)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v87
	v111 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v111
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+25)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v93)+56)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v93)+40)) = v111
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+118)))
	switch v121 - int32(112) {
	case 0, 5:
		goto L20
	default:
		goto L21
	case 4:
		goto L22
	}
L19:
	;
	v176 = v93 + int32(56)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+72))
	if v179 != 0 {
		goto L40
	} else {
		goto L41
	}
L20:
	;
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v169)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(-1)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L37
	}
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+68))
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[5]))
	if v128 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	if v136 != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	v129 = int32(1)
	if v124 == v128 {
		v136 = v129
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v136 = int32(0)
	goto L24
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[6]))
	if v132 == v124 {
		v136 = v129
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[7]))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[8]))
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v141)
	if v140 == int32(-1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+68))
	v149 = F_GetTempNamespaceProcNumber(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L36
	}
L33:
	;
	v145 = v138
	goto L35
L34:
	;
	v145 = v140
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v145
	goto L19
L36:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+24)) = uint8(v151)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = v149
	goto L19
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v158)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v159
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_2), v22)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(1197), int32(_a_F_RelationBuildDesc_4))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v181 = v179
	goto L42
L41:
	;
	v181 = int32(2249)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = int32(-1)
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v190 = F_MemoryContextAllocZero(m, v188, int32(20))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v193 = v22 + int32(160)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	F_ScanKeyInit(m, v193, int32(1), int32(3), int32(184), v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v200 = int32(5)
	F_ScanKeyInit(m, v22+int32(208), v200, v200, int32(146), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v208 = F_table_open(m, int32(1249), int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationBuildDesc[9])))
	v215 = F_systable_beginscan(m, v208, int32(2659), v212, int32(0), int32(2), v193)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217)+120)))
	v224 = v218
	v226 = int32(0)
	v231 = int32(0)
	goto L49
L48:
	;
	F_systable_endscan(m, v215)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L85
	}
L49:
	;
	v239 = F_systable_getnext(m, v215)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	v373 = int32(0)
	v374 = v368
	v378 = v367
	goto L48
L51:
	;
	if v239 == int32(0) {
		v373 = v224
		v374 = v226
		v378 = v231
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+22)))
	v245 = v243 + v244
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+74)))
	if v246 <= int32(0) {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+120)))
	if v250 < v246 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v260 = (v246 - int32(1)) & int32(_a_F_RelationBuildDesc_5)
	v261 = int32(100)
	base.MemoryCopy(m, v252+v253<<(uint(int32(4))%32)+v260*v261+int32(20), v245, v261)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	F_populate_compact_attribute(m, v268, v260)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+86)))
	if v271 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+16)) = uint8(v274)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+90)))
	if v276 == int32(115) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v279 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+17)) = uint8(v279)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+90)))
	v282 = v281
	goto L61
L60:
	;
	v282 = v276
	goto L61
L61:
	;
	if v282&int32(255) == int32(118) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v287 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v190)+18)) = uint8(v287)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+87)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+88)))
	if v290 != int32(1) {
		v367 = v231
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v368 = v226 + v289
	v370 = v224 - int32(1)
	if v370 != 0 {
		v224 = v370
		v226 = v368
		v231 = v367
		goto L49
	} else {
		goto L84
	}
L66:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v208)+52))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v239)+16))
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v294)+18)))
	if base.Ui32(v295&int32(2047)) <= base.Ui32(int32(24)) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v311 != 0 {
		v367 = v231
		goto L65
	} else {
		goto L73
	}
L68:
	;
	v303 = F_getmissingattr(m, v293, int32(25), v22+int32(271))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v308 = F_fastgetattr_3(m, v239, int32(25), v293, v22+int32(271))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L72
	}
L71:
	;
	v310 = v303
	goto L67
L72:
	;
	v310 = v308
	goto L67
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(1)
	if v231 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v318)+120)))
	v322 = F_MemoryContextAllocZero(m, v317, v319<<(uint(int32(3))%32))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	v324 = v231
	goto L76
L76:
	;
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+72)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+82)))
	v331 = int32(*(*int8)(unsafe.Add(mBase, uint32(v245)+83)))
	v334 = F_array_get_element(m, v310, int32(1), v22+int32(272), int32(-1), v329, v330, v331, v22+int32(159))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L78
	}
L77:
	;
	v324 = v322
	goto L76
L78:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+82)))
	if v336 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v363 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v324+v260<<(uint(int32(3))%32)))) = uint8(v363)
	v367 = v324
	goto L65
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v260<<(uint(int32(3))%32))+4)) = v334
	goto L79
L81:
	;
	goto L82
L82:
	;
	v343 = int32(_a_F_RelationBuildDesc_1)
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3]))
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v347
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+82)))
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+72)))
	v354 = F_datumCopy(m, v334, v352, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v324+v260<<(uint(int32(3))%32))+4)) = v354
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[3])) = v344
	goto L79
L84:
	;
	goto L50
L85:
	;
	F_relation_close(m, v208, int32(1))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	if v373 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v386 = int32(*(*int16)(unsafe.Add(mBase, uint32(v385)+120)))
	if int32(0) < v386 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+20)) = int32(0)
	goto L90
L89:
	;
	goto L90
L90:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+16)))
	if v392 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v936 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+128)) = v936
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+88)) = uint8(v936)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+84)) = v936
	v942 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+92)) = v942
	*(*int64)(unsafe.Add(mBase, uint32(v93)+100)) = v942
	*(*int64)(unsafe.Add(mBase, uint32(v93)+108)) = v942
	*(*int64)(unsafe.Add(mBase, uint32(v93)+116)) = v942
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+124)) = uint8(v936)
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v952)+119)))
	switch v953 - int32(73) {
	case 0, 32:
		goto L212
	default:
		goto L210
	case 10, 36, 41, 43:
		goto L211
	}
L92:
	;
	F_pfree(m, v190)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L4
	} else {
		goto L209
	}
L93:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	v405 = base.B2i32(base.Ui32(v403) < base.Ui32(int32(_a_F_RelationBuildDesc_6)))
	goto L98
L94:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+17)))
	if v393 != 0 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+18)))
	if v378|(v394|base.B2i32(int32(0) < v374)) != 0 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v400 = int32(*(*int16)(unsafe.Add(mBase, uint32(v399)+122)))
	if v400 <= int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	goto L93
L98:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v406)+16)) = v190
	if int32(0) < v374 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v378
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v618)+122)))
	if v619 <= int32(0) {
		goto L148
	} else {
		goto L149
	}
L100:
	;
	v410 = int32(0)
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v415 = F_MemoryContextAllocZero(m, v412, v374<<(uint(int32(3))%32))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v596 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v190)+12)) = uint16(v596)
	goto L99
L103:
	;
	v418 = v22 + int32(272)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	F_ScanKeyInit(m, v418, int32(2), int32(3), int32(184), v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v427 = F_table_open(m, int32(2604), int32(1))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L106
	}
L105:
	;
	F_systable_endscan(m, v433)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L133
	}
L106:
	;
	v430 = int32(1)
	v433 = F_systable_beginscan(m, v427, int32(2656), v430, int32(0), v430, v418)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v435 = F_systable_getnext(m, v433)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	if v435 == int32(0) {
		v539 = v410
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v442 = v410
	v443 = v435
	goto L110
L110:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v443)+16))
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+22)))
	v460 = v458 + v459
	if v374 <= v442 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v539 = v531
	goto L105
L112:
	;
	v464 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v427)+52))
	v488 = F_fastgetattr_3(m, v443, int32(4), v485, v22+int32(271))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L119
	}
L115:
	;
	if v464 == int32(0) {
		v539 = v442
		goto L105
	} else {
		goto L116
	}
L116:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v469 = int32(*(*int16)(unsafe.Add(mBase, uint32(v460)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v468 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_7), v22+int32(112))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_8), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	v539 = v442
	goto L105
L119:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v490 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v534 = F_systable_getnext(m, v433)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L131
	}
L121:
	;
	v495 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v515 = F_text_to_cstring(m, v488)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L4
	} else {
		goto L128
	}
L124:
	;
	if v495 == int32(0) {
		v531 = v442
		goto L120
	} else {
		goto L125
	}
L125:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v500 = int32(*(*int16)(unsafe.Add(mBase, uint32(v460)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v500
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v499 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_10), v22+int32(96))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_11), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v531 = v442
	goto L120
L128:
	;
	v519 = v415 + v442<<(uint(int32(3))%32)
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v519))) = uint16(v520)
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v524 = F_MemoryContextStrdup(m, v523, v515)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v524
	F_pfree(m, v515)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v531 = v442 + int32(1)
	goto L120
L131:
	;
	if v534 != 0 {
		v442 = v531
		v443 = v534
		goto L110
	} else {
		goto L132
	}
L132:
	;
	goto L111
L133:
	;
	F_relation_close(m, v427, int32(1))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	if v539 == v374 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	if int32(2) <= v539 {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v563 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	if v563 == int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v374 - v539
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v567 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_12), v22+int32(80))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_13), int32(_a_F_RelationBuildDesc_9))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	goto L135
L141:
	;
	F_pg_qsort(m, v415, v539, int32(8), int32(1590))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v590)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v591))) = v415
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v594)+12)) = uint16(v539)
	goto L99
L144:
	;
	goto L143
L145:
	;
	v909 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v893)+122)))
	if v909 != 0 {
		goto L91
	} else {
		goto L208
	}
L146:
	;
	v852 = int32(0)
	v853 = int32(*(*int16)(unsafe.Add(mBase, uint32(v836)+120)))
	if v853 <= v852 {
		v893 = v836
		goto L145
	} else {
		goto L201
	}
L147:
	;
	v634 = v22 + int32(272)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	F_ScanKeyInit(m, v634, int32(9), int32(3), int32(184), v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L4
	} else {
		goto L154
	}
L148:
	;
	if base.Ui32(v403) < base.Ui32(int32(_a_F_RelationBuildDesc_6)) {
		v893 = v618
		goto L145
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v630 = F_MemoryContextAllocZero(m, v627, v619*int32(12))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L4
	} else {
		goto L153
	}
L151:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+16)))
	if v622 != int32(1) {
		v836 = v618
		goto L146
	} else {
		goto L152
	}
L152:
	;
	v632 = int32(0)
	goto L147
L153:
	;
	v632 = v630
	goto L147
L154:
	;
	v641 = int32(0)
	v644 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L156
	}
L155:
	;
	F_systable_endscan(m, v650)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L188
	}
L156:
	;
	v647 = int32(1)
	v650 = F_systable_beginscan(m, v644, int32(2665), v647, int32(0), v647, v634)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L157
	}
L157:
	;
	v652 = F_systable_getnext(m, v650)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	if v652 == int32(0) {
		v778 = v641
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v659 = v652
	v662 = v641
	goto L160
L160:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v659)+16))
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675)+22)))
	v677 = v675 + v676
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+72)))
	switch v678 - int32(99) {
	case 0:
		goto L163
	default:
		v768 = v662
		goto L162
	case 11:
		goto L164
	}
L161:
	;
	v778 = v768
	goto L155
L162:
	;
	v770 = F_systable_getnext(m, v650)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L186
	}
L163:
	;
	if v619 <= v662 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+76)))
	if v681 != 0 {
		v768 = v662
		goto L162
	} else {
		goto L165
	}
L165:
	;
	v682 = F_extractNotNullColumn(m, v659)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v688 = int32(105)
	*(*uint8)(unsafe.Add(mBase, uint32(v684+v682<<(uint(int32(4))%32))+15)) = uint8(v688)
	v768 = v662
	goto L162
L167:
	;
	v693 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L4
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v713 = v632 + v662*int32(12)
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v713)+8)) = uint8(v714)
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v713)+9)) = uint8(v716)
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v713)+10)) = uint8(v718)
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v724 = F_MemoryContextStrdup(m, v721, v677+int32(4))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L4
	} else {
		goto L174
	}
L170:
	;
	if v693 == int32(0) {
		v778 = v662
		goto L155
	} else {
		goto L171
	}
L171:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v697 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_14), v22-int32(-64))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_15), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	v778 = v662
	goto L155
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = v724
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v644)+52))
	v731 = F_fastgetattr_3(m, v659, int32(28), v728, v22+int32(271))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+271)))
	if v733 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v738 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v756 = F_text_to_cstring(m, v731)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L4
	} else {
		goto L183
	}
L179:
	;
	if v738 == int32(0) {
		v768 = v662
		goto L162
	} else {
		goto L180
	}
L180:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v742 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_17), v22+int32(48))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_18), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v768 = v662
	goto L162
L183:
	;
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[4]))
	v760 = F_MemoryContextStrdup(m, v759, v756)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713)+4)) = v760
	F_pfree(m, v756)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v768 = v662 + int32(1)
	goto L162
L186:
	;
	if v770 != 0 {
		v659 = v770
		v662 = v768
		goto L160
	} else {
		goto L187
	}
L187:
	;
	goto L161
L188:
	;
	F_relation_close(m, v644, int32(1))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L4
	} else {
		goto L189
	}
L189:
	;
	if v778 == v619 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	if int32(2) <= v778 {
		goto L196
	} else {
		goto L197
	}
L191:
	;
	v799 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	if v799 == int32(0) {
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v619 - v778
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v803 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_19), v22+int32(32))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L4
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(_a_F_RelationBuildDesc_20), int32(_a_F_RelationBuildDesc_16))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	goto L190
L196:
	;
	F_pg_qsort(m, v632, v778, int32(12), int32(1591))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v827)+4)) = v632
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v830)+14)) = uint16(v778)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	if base.Ui32(v403) < base.Ui32(int32(_a_F_RelationBuildDesc_6)) {
		v893 = v832
		goto L145
	} else {
		goto L200
	}
L199:
	;
	goto L198
L200:
	;
	v836 = v832
	goto L146
L201:
	;
	v859 = v836
	v860 = v852
	goto L202
L202:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	v878 = v875 + v860<<(uint(int32(4))%32)
	v879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878)+31)))
	if v879 == int32(117) {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v893 = v885
	goto L145
L204:
	;
	v882 = int32(118)
	*(*uint8)(unsafe.Add(mBase, uint32(v878)+31)) = uint8(v882)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v885 = v884
	goto L206
L205:
	;
	v885 = v859
	goto L206
L206:
	;
	v887 = v860 + int32(1)
	v888 = int32(*(*int16)(unsafe.Add(mBase, uint32(v885)+120)))
	if v887 < v888 {
		v859 = v885
		v860 = v887
		goto L202
	} else {
		goto L207
	}
L207:
	;
	goto L203
L208:
	;
	v910 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v190)+14)) = uint16(v910)
	goto L91
L209:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v93)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v914)+16)) = int32(0)
	goto L91
L210:
	;
	F_RelationParseRelOptions(m, v93, v77)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L4
	} else {
		goto L215
	}
L211:
	;
	F_RelationInitTableAccessMethod(m, v93)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L214
	}
L212:
	;
	F_RelationInitIndexAccessInfo(m, v93)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	goto L210
L214:
	;
	goto L210
L215:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962)+124)))
	if v963 == int32(1) {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+125)))
	if v972 == int32(1) {
		goto L222
	} else {
		goto L223
	}
L217:
	;
	F_RelationBuildRuleLock(m, v93)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L4
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93)+68)) = int64(0)
	v971 = v962
	goto L216
L220:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v971 = v968
	goto L216
L221:
	;
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+127)))
	if v981 == int32(1) {
		goto L227
	} else {
		goto L228
	}
L222:
	;
	F_RelationBuildTriggers(m, v93)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L4
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+76)) = int32(0)
	v980 = v971
	goto L221
L225:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v980 = v977
	goto L221
L226:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v93)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+60)) = v988
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[10]))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993)+117)))
	if v994 != 0 {
		goto L232
	} else {
		goto L233
	}
L227:
	;
	F_RelationBuildRowSecurity(m, v93)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L4
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+80)) = int32(0)
	goto L226
L230:
	;
	goto L226
L231:
	;
	F_RelationInitPhysicalAddr(m, v93)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L235
	}
L232:
	;
	v995 = int32(0)
	goto L234
L233:
	;
	v995 = v992
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+64)) = v995
	goto L231
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = int32(0)
	F_pfree(m, v77)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004+v52)+4)))
	if v1006 != int32(1) {
		goto L9
	} else {
		goto L237
	}
L237:
	;
	F_RelationDestroyRelation(m, v93, int32(0))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[0]))
	v1015 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1013+v52)+4)) = uint8(v1015)
	v1019 = F_ScanPgRelation(m, l0, int32(1), v1015)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	if v1019 != 0 {
		v77 = v1019
		goto L14
	} else {
		goto L240
	}
L240:
	;
	goto L15
L241:
	;
	v1100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+26)) = uint8(v1100)
	v1104 = v93
	goto L8
L242:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[11]))
	v1060 = F_hash_search(m, v1056, v176, int32(1), v22+int32(160))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+160)))
	if v1062 == int32(1) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1060)+4)) = v93
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+16))
	if v1067 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1060)+4)) = v93
	goto L241
L247:
	;
	F_RelationDestroyRelation(m, v1065, int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L4
	} else {
		goto L250
	}
L248:
	;
	goto L249
L249:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildDesc[12]))
	if v1074 == int32(0) {
		goto L241
	} else {
		goto L251
	}
L250:
	;
	goto L241
L251:
	;
	v1079 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	if v1079 == int32(0) {
		goto L241
	} else {
		goto L253
	}
L253:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1083 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_21), v22+int32(16))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(1314), int32(_a_F_RelationBuildDesc_4))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	goto L241
L256:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+48))
	v1130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v1130
	*(*int32)(unsafe.Add(mBase, uint32(v22)+148)) = v1129 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_22), v22+int32(144))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(585), int32(_a_F_RelationBuildDesc_23))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v1149
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v373
	F_errmsg_internal(m, int32(_a_F_RelationBuildDesc_24), v22+int32(128))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	F_errfinish(m, int32(_a_F_RelationBuildDesc_3), int32(667), int32(_a_F_RelationBuildDesc_23))
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L4
	} else {
		goto L261
	}
L261:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCacheInitFilePostInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePostInvalidate[0]))
	F_LWLockRelease(m, v2+int32(2048))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_RelationCacheInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
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
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	v1 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[0]))
	if v13 == int32(_a_F_RelationCacheInvalidate_0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_read_relmap_file(m, int32(_a_F_RelationCacheInvalidate_1), int32(_a_F_RelationCacheInvalidate_2), int32(0), int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[1]))
	if v23 == int32(_a_F_RelationCacheInvalidate_0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[2]))
	F_read_relmap_file(m, int32(_a_F_RelationCacheInvalidate_3), v28, int32(0), int32(22))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v34 = v10 + int32(12)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[3]))
	F_hash_seq_init(m, v34, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v39 = F_hash_seq_search(m, v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L14
	}
L11:
	;
	F_list_free(m, v311)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L127
	}
L12:
	;
	F_list_free(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L94
	}
L13:
	;
	v231 = v128
	goto L12
L14:
	;
	if v39 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = v39
	v44 = v1
	v45 = v1
	goto L18
L16:
	;
	goto L17
L17:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L92
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+32))
	if v49 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_smgrreleaseall(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L61
	}
L20:
	;
	v131 = F_hash_seq_search(m, v10+int32(12))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L59
	}
L21:
	;
	v55 = int32(_a_F_RelationCacheInvalidate_4)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[4])) = v57 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v61 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
	if v52 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v127 = v44
	v128 = v45
	goto L20
L25:
	;
	goto L24
L26:
	;
	v127 = v124
	v128 = v125
	goto L20
L27:
	;
	F_RelationClearRelation(m, v48)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	switch v67 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L32
	default:
		goto L31
	}
L30:
	;
	v124 = v44
	v125 = v45
	goto L26
L31:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
	if v107 != int32(2662) {
		goto L48
	} else {
		goto L49
	}
L32:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+88))
	if v70 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v71 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	v76 = v74 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v76
	if v76 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	F_RelationInitPhysicalAddr(m, v48)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L46
	}
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	F_smgrclose(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L45
	}
L38:
	;
	v81 = v71 + int32(76)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v83 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v90
	v92 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+80)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v81
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v81
	goto L40
L42:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v90 = v85
	goto L41
L43:
	;
	goto L44
L44:
	;
	v87 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v87
	v90 = v87
	goto L41
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = int32(0)
	goto L36
L46:
	;
	goto L31
L47:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+25)))
	if v116 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	if v107 != int32(1259) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v114 = F_lappend(m, v45, v48)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	v112 = F_lcons(m, v48, v45)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v124 = v44
	v125 = v112
	goto L26
L53:
	;
	v124 = v44
	v125 = v114
	goto L26
L54:
	;
	v119 = F_lcons(m, v48, v44)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v121 = F_lappend(m, v44, v48)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L58
	}
L57:
	;
	v124 = v119
	v125 = v45
	goto L26
L58:
	;
	v124 = v121
	v125 = v45
	goto L26
L59:
	;
	if v131 != 0 {
		v41 = v131
		v44 = v127
		v45 = v128
		goto L18
	} else {
		goto L60
	}
L60:
	;
	goto L19
L61:
	;
	v135 = int32(0)
	if v128 == v135 {
		v231 = v135
		goto L12
	} else {
		goto L62
	}
L62:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v138 <= int32(0) {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	v144 = int32(0)
	goto L64
L64:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v150 = int32(2)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v144<<(uint(v150)%32))))
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[7]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+20))
	goto L68
L65:
	;
	goto L13
L66:
	;
	v209 = v144 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v209 < v210 {
		v144 = v209
		goto L64
	} else {
		goto L91
	}
L67:
	;
	F_RelationRebuildRelation(m, v153)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L90
	}
L68:
	;
	if v156 == v150 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+25)))
	if v159 != int32(1) {
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	if v165 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	if v162 != int32(1) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+72))
	v170 = v168 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+72)) = v170
	if v170 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v153)+256))
	if v198 != 0 {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	F_smgrclose(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L85
	}
L78:
	;
	v175 = v165 + int32(76)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v177 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+76)) = v184
	v186 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+80)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v184)+4)) = v175
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v175
	goto L80
L82:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v184 = v179
	goto L81
L83:
	;
	goto L84
L84:
	;
	v181 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v181
	v184 = v181
	goto L81
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(0)
	goto L76
L86:
	;
	F_pfree(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+26)) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+256)) = v201
	goto L66
L89:
	;
	goto L88
L90:
	;
	goto L66
L91:
	;
	goto L65
L92:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v311 = v1
	goto L11
L94:
	;
	if v127 == int32(0) {
		v311 = v1
		goto L11
	} else {
		goto L95
	}
L95:
	;
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v237 <= v236 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v311 = v127
	goto L11
L97:
	;
	goto L98
L98:
	;
	v242 = v236
	goto L99
L99:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v248 = int32(2)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v242<<(uint(v248)%32))))
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[7]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	goto L103
L100:
	;
	v311 = v127
	goto L11
L101:
	;
	v307 = v242 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v307 < v308 {
		v242 = v307
		goto L99
	} else {
		goto L126
	}
L102:
	;
	F_RelationRebuildRelation(m, v251)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L125
	}
L103:
	;
	if v254 == v248 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+25)))
	if v257 != int32(1) {
		goto L102
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	if v263 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v260 != int32(1) {
		goto L102
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+72))
	v268 = v266 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+72)) = v268
	if v268 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	goto L111
L111:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v251)+256))
	if v296 != 0 {
		goto L121
	} else {
		goto L122
	}
L112:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	F_smgrclose(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L120
	}
L113:
	;
	v273 = v263 + int32(76)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5]))
	if v275 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	goto L112
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+76)) = v282
	v284 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+80)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v282)+4)) = v273
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6])) = v273
	goto L115
L117:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[6]))
	v282 = v277
	goto L116
L118:
	;
	goto L119
L119:
	;
	v279 = int32(_a_F_RelationCacheInvalidate_5)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[5])) = v279
	v282 = v279
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+12)) = int32(0)
	goto L111
L121:
	;
	F_pfree(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v251)+26)) = uint8(v299)
	*(*int32)(unsafe.Add(mBase, uint32(v251)+256)) = v299
	goto L101
L124:
	;
	goto L123
L125:
	;
	goto L101
L126:
	;
	goto L100
L127:
	;
	v319 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[8]))
	if v321 <= v319 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	m.G0 = v10 + int32(32)
	return
L129:
	;
	v325 = v321 & int32(7)
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInvalidate[9]))
	if base.Ui32(int32(8)) <= base.Ui32(v321) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v333 = v319
	v336 = int32(0)
	goto L133
L131:
	;
	v366 = v319
	goto L132
L132:
	;
	v374 = v366
	v376 = int32(0)
	goto L137
L133:
	;
	v342 = v327 + v333<<(uint(int32(3))%32)
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+60)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+52)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+44)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+36)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+28)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+20)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+12)) = uint8(v343)
	*(*uint8)(unsafe.Add(mBase, uint32(v342)+4)) = uint8(v343)
	v359 = int32(8)
	v360 = v333 + v359
	v362 = v336 + v359
	if v362 != v321&int32(2147483640) {
		v333 = v360
		v336 = v362
		goto L133
	} else {
		goto L135
	}
L134:
	;
	if v325 == int32(0) {
		goto L128
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	v366 = v360
	goto L132
L137:
	;
	v384 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v374<<(uint(int32(3))%32))+4)) = uint8(v384)
	v389 = v376 + v384
	if v389 != v325 {
		v374 = v374 + v384
		v376 = v389
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L128
L139:
	;
	goto L138
}
func F_RelationClearMissing(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int64
	_ = v20
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+120)))
	base.MemoryFill(m, v10+int32(80), int32(0), int32(100))
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v20
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v32)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+72)) = uint8(v32)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v32)
	v40 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < v14 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v45 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_relation_close(m, v40, int32(3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L19
	}
L7:
	;
	v54 = F_SearchSysCache2(m, int32(7), v12, base.I32_extend16_s(v45))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v54 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v59)+88)))
	if v61 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v40)+52))
	v71 = F_heap_modify_tuple(m, v54, v64, v10+int32(80), v10+int32(48), v10+int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseCatCache(m, v54)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L17
	}
L14:
	;
	F_CatalogTupleUpdate(m, v40, v71+int32(4), v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_pfree(m, v71)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v45 != v14 {
		v45 = v45 + int32(1)
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L8
L19:
	;
	m.G0 = v10 + int32(192)
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v45
	F_errmsg_internal(m, int32(_a_F_RelationClearMissing_0), v10)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_RelationClearMissing_1), int32(1999), int32(_a_F_RelationClearMissing_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationForgetRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[0]))
	v15 = F_hash_search(m, v10, v6+int32(12), v2, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 == int32(0) {
			m.G0 = v6 + int32(16)
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v19 == int32(0) {
				m.G0 = v6 + int32(16)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				if v22 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v84
						F_errmsg_internal(m, int32(_a_F_RelationForgetRelation_0), v6)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationForgetRelation_1), int32(2903), int32(_a_F_RelationForgetRelation_2))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
					if v23 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
						if v26 == int32(0) {
							F_RelationClearRelation(m, v19)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[1]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							if v33 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
								v38 = v36 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
								if v38 == int32(0) {
									v43 = v33 + int32(76)
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2]))
									if v45 != 0 {
										v47 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3]))
										v52 = v47
									} else {
										v49 = int32(_a_F_RelationForgetRelation_3)
										*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2])) = v49
										v52 = v49
									}
									*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
									v54 = int32(_a_F_RelationForgetRelation_3)
									*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
									*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
									*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3])) = v43
								} else {
								}
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
								F_smgrclose(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
									if v66 != 0 {
										F_pfree(m, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v69 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
											*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
											m.G0 = v6 + int32(16)
											return
										}
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[1]))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v31
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if v33 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+72))
							v38 = v36 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v33)+72)) = v38
							if v38 == int32(0) {
								v43 = v33 + int32(76)
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2]))
								if v45 != 0 {
									v47 = *(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3]))
									v52 = v47
								} else {
									v49 = int32(_a_F_RelationForgetRelation_3)
									*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[2])) = v49
									v52 = v49
								}
								*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v52
								v54 = int32(_a_F_RelationForgetRelation_3)
								*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_RelationForgetRelation[3])) = v43
							} else {
							}
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
							F_smgrclose(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
								if v66 != 0 {
									F_pfree(m, v66)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										v69 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
										*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v19)+256))
							if v66 != 0 {
								F_pfree(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									v69 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v69 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+26)) = uint8(v69)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+256)) = v69
								m.G0 = v6 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetFKeyList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v13 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v116
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v116 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v11, int32(9), int32(3), int32(184), v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v27 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v30 = int32(1)
	v33 = F_systable_beginscan(m, v27, int32(2665), v30, int32(0), v30, v11)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v35 = F_systable_getnext(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v2
	v41 = v35
	goto L13
L11:
	;
	v87 = v2
	goto L12
L12:
	;
	F_systable_endscan(m, v33)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L23
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v47 = v45 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+72)))
	if v48 == int32(102) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v87 = v81
	goto L12
L15:
	;
	v52 = F_palloc0(m, int32(280))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	v81 = v40
	goto L17
L17:
	;
	v82 = F_systable_getnext(m, v33)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(470)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+75)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+20)) = uint8(v62)
	v72 = int32(0)
	F_DeconstructFkConstraintRow(m, v41, v52+int32(16), v52+int32(22), v52+int32(86), v52+int32(152), v72, v72, v72, v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v78 = F_lappend(m, v40, v52)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v81 = v78
	goto L17
L21:
	;
	if v82 != 0 {
		v40 = v81
		v41 = v82
		goto L13
	} else {
		goto L22
	}
L22:
	;
	goto L14
L23:
	;
	F_relation_close(m, v27, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v97 = int32(_a_F_RelationGetFKeyList_0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0]))
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0])) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v104 = F_copyObjectImpl(m, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v106 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v106)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v104
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetFKeyList[0])) = v98
	F_list_free_deep(m, v103)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	v116 = v87
	goto L1
}
func F_RelationGetIndexScan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	v7 = F_palloc(m, int32(92))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		if v11 < l1 {
			v24 = F_palloc(m, l1*int32(48))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v27 = v24
				*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v27
				if int32(0) < l2 {
					v33 = F_palloc(m, l2*int32(48))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v33
						v36 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
						v44 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
						*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
						*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
						v52 = v43 ^ int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
						return v7
					}
				} else {
					v35 = int32(0)
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
					v52 = v43 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
					return v7
				}
			}
		} else {
			v27 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v27
			if int32(0) < l2 {
				v33 = F_palloc(m, l2*int32(48))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = v33
					v36 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
					v44 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
					v52 = v43 ^ int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
					return v7
				}
			} else {
				v35 = int32(0)
				v36 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)) = uint8(v36)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v36)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v35
				v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexScan[0]))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+69)))
				v44 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+36)) = v44
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)) = uint8(v43)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+44)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v7)+52)) = v44
				v52 = v43 ^ int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)) = uint8(v52)
				return v7
			}
		}
	}
}
func F_RelationIncrementReferenceCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[0]))
	F_ResourceOwnerEnlarge(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6 + int32(1)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[1]))
		if v11 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIncrementReferenceCount[0]))
			F_ResourceOwnerRemember(m, v13, l0, int32(_a_F_RelationIncrementReferenceCount_0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_RelationInitLockInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v2
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitLockInfo[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+117)))
	if v8 != 0 {
		v9 = int32(0)
	} else {
		v9 = v6
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v9
	return
}
func F_RelationIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v133
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v20 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v20)
	v133 = v3
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_RelationIsVisibleExt_0), v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_RelationIsVisibleExt_1), int32(940), int32(_a_F_RelationIsVisibleExt_2))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v39 = v35 + v36
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v40 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v14)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L40
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIsVisibleExt[0]))
	v45 = int32(0)
	if v44 == v45 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIsVisibleExt[0]))
	if v87 == int32(0) {
		v123 = v3
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v83 == int32(0) {
		v123 = v3
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v83 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v51 <= int32(0) {
		v77 = v45
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v83 = v77
	goto L18
L23:
	;
	v54 = int32(0)
	if v54 < v51 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v57 = v51
	goto L26
L25:
	;
	v57 = v54
	goto L26
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v60 = int32(0)
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58+v60<<(uint(int32(2))%32))))
	v69 = base.B2i32(v68 == v40)
	if v68 == v40 {
		v77 = v69
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v77 = v69
	goto L22
L29:
	;
	v71 = v60 + int32(1)
	if v71 != v57 {
		v60 = v71
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= v90 {
		v123 = v3
		goto L14
	} else {
		goto L33
	}
L33:
	;
	v96 = v90
	goto L34
L34:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v96<<(uint(int32(2))%32))))
	v109 = base.B2i32(v108 == v40)
	if v108 == v40 {
		v123 = v109
		goto L14
	} else {
		goto L36
	}
L35:
	;
	v123 = v109
	goto L14
L36:
	;
	v110 = F_get_relname_relid(m, v39+int32(4), v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v110 != 0 {
		v123 = v109
		goto L14
	} else {
		goto L38
	}
L38:
	;
	v113 = v96 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v113 < v114 {
		v96 = v113
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v133 = v123
	goto L1
}
func F_RelationMapInvalidate(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[0]))
		if v3 != int32(_a_F_RelationMapInvalidate_0) {
			return
		} else {
			F_read_relmap_file(m, int32(_a_F_RelationMapInvalidate_1), int32(_a_F_RelationMapInvalidate_2), int32(0), int32(22))
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[1]))
		if v13 != int32(_a_F_RelationMapInvalidate_0) {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapInvalidate[2]))
			F_read_relmap_file(m, int32(_a_F_RelationMapInvalidate_3), v18, int32(0), int32(22))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_relation_needs_vacanalyze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 float64
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v116 float64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 float32
	_ = v184
	var v185 float32
	_ = v185
	var v188 float32
	_ = v188
	var v189 float32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v204 float32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v212 float32
	_ = v212
	var v215 float32
	_ = v215
	var v219 float32
	_ = v219
	var v221 float32
	_ = v221
	var v222 float32
	_ = v222
	var v223 float32
	_ = v223
	var v226 float32
	_ = v226
	var v230 float32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v257 int32
	_ = v257
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	v30 = m.G0
	v32 = v30 - int32(112)
	m.G0 = v32
	if l1 != 0 {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
		if v34 < l4 {
			v36 = v34
		} else {
			v36 = l4
		}
		if v34 < int32(0) {
			v39 = l4
		} else {
			v39 = v36
		}
		v41 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[0]))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		if v42 < v41 {
			v44 = v42
		} else {
			v44 = v41
		}
		if v42 < int32(0) {
			v47 = v41
		} else {
			v47 = v44
		}
		v49 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[1]))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if v50 < int32(0) {
			v53 = v49
		} else {
			v53 = v50
		}
		v54 = *(*float64)(unsafe.Add(mBase, uint32(l1)+80))
		v56 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[2]))
		if base.F64_ge(v54, float64(0)) != 0 {
			v59 = v54
		} else {
			v59 = v56
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v62 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[3]))
		if int32(-2) < v60 {
			v65 = v60
		} else {
			v65 = v62
		}
		v66 = *(*float64)(unsafe.Add(mBase, uint32(l1)+72))
		v68 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[4]))
		if base.F64_ge(v66, float64(0)) != 0 {
			v71 = v66
		} else {
			v71 = v68
		}
		v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v74 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[5]))
		if int32(-2) < v72 {
			v77 = v72
		} else {
			v77 = v74
		}
		v79 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[6]))
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if v80 < int32(0) {
			v83 = v79
		} else {
			v83 = v80
		}
		v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+64))
		v86 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[7]))
		if base.F64_ge(v84, float64(0)) != 0 {
			v89 = v84
		} else {
			v89 = v86
		}
		v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v108 = v39
		v109 = v65
		v110 = v77
		v111 = v83
		v112 = v47
		v113 = v53
		v114 = v59
		v115 = v71
		v116 = v89
		v117 = v90
	} else {
		v92 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[0]))
		v94 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[1]))
		v96 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[2]))
		v98 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[3]))
		v100 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[4]))
		v102 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[5]))
		v104 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[6]))
		v106 = *(*float64)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[7]))
		v108 = l4
		v109 = v98
		v110 = v102
		v111 = v104
		v112 = v92
		v113 = v94
		v114 = v96
		v115 = v100
		v116 = v106
		v117 = int32(1)
	}
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+136))
	if base.Ui32(v118) < base.Ui32(int32(3)) {
		v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
		if v147 != 0 {
			v150 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[8]))
			if v150 == v108 {
				v153 = int32(-1)
			} else {
				v153 = v150 - v108
			}
			v159 = int32(base.Ui32(v147-v153) >> (uint(int32(31)) % 32))
		} else {
			v159 = int32(0)
		}
		v160 = int32(1)
		v161 = v159 & v160
		*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v161)
		if (v159|v117)&v160 != 0 {
			v170 = v159
			if l3 == int32(0) {
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
				v308 = int32(0)
				if l0 != int32(2619) {
					v312 = v308
				} else {
					v312 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
				m.G0 = v32 + int32(112)
				return
			} else {
				v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[9])))
				if v174 != int32(1) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
					v308 = int32(0)
					if l0 != int32(2619) {
						v312 = v308
					} else {
						v312 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
					m.G0 = v32 + int32(112)
					return
				} else {
					v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[10])))
					if v178&int32(1) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
						v308 = int32(0)
						if l0 != int32(2619) {
							v312 = v308
						} else {
							v312 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
						m.G0 = v32 + int32(112)
						return
					} else {
						v184 = float32(0)
						v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
						if base.F32_lt(v185, v184) != 0 {
							v188 = v184
						} else {
							v188 = v185
						}
						v189 = float32(1)
						v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
						if v190 <= int32(0) {
							v204 = v189
						} else {
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
							if v193 <= int32(0) {
								v204 = v189
							} else {
								if base.Ui32(v193) < base.Ui32(v190) {
									v198 = v193
								} else {
									v198 = v190
								}
								v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
							}
						}
						v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
						v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
						v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
						v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v111))
						if v110 < int32(0) {
							v219 = v212
						} else {
							v215 = base.F32_convert_i32_u(v110)
							if base.F32_gt(v212, v215) == int32(0) {
								v219 = v212
							} else {
								v219 = v215
							}
						}
						v221 = base.F32_convert_i64_s(v205)
						v222 = base.F32_convert_i64_s(v206)
						v223 = base.F32_convert_i64_s(v207)
						v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v114), v188), base.F32_convert_i32_s(v113))
						v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v115)), v204), base.F32_convert_i32_s(v109))
						v233 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return
						} else {
							if int32(0) <= v109 {
								if v233 == int32(0) {
									if v170|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v308 = base.F32_gt(v221, v226)
									if l0 != int32(2619) {
										v312 = v308
									} else {
										v312 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
									F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_0), v32)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										v278 = int32(3138)
										F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v170|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v308 = base.F32_gt(v221, v226)
											if l0 != int32(2619) {
												v312 = v308
											} else {
												v312 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							} else {
								if v233 == int32(0) {
									if v170|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v308 = base.F32_gt(v221, v226)
									if l0 != int32(2619) {
										v312 = v308
									} else {
										v312 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
									F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_3), v32-int32(-64))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return
									} else {
										v278 = int32(3142)
										F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v170|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v308 = base.F32_gt(v221, v226)
											if l0 != int32(2619) {
												v312 = v308
											} else {
												v312 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v166 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v166)
			*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v166)
			m.G0 = v32 + int32(112)
			return
		}
	} else {
		v122 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[11]))
		v123 = v122 - v112
		v124 = int32(3)
		if base.Ui32(v123) < base.Ui32(v124) {
			v128 = v123 - v124
		} else {
			v128 = v123
		}
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v128))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v118)) == int32(0) {
			v140 = base.B2i32(base.Ui32(v118) < base.Ui32(v128))
		} else {
			v140 = int32(base.Ui32(v118-v128) >> (uint(int32(31)) % 32))
		}
		if v140 == int32(0) {
			v147 = *(*int32)(unsafe.Add(mBase, uint32(l2)+140))
			if v147 != 0 {
				v150 = *(*int32)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[8]))
				if v150 == v108 {
					v153 = int32(-1)
				} else {
					v153 = v150 - v108
				}
				v159 = int32(base.Ui32(v147-v153) >> (uint(int32(31)) % 32))
			} else {
				v159 = int32(0)
			}
			v160 = int32(1)
			v161 = v159 & v160
			*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v161)
			if (v159|v117)&v160 != 0 {
				v170 = v159
				if l3 == int32(0) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
					v308 = int32(0)
					if l0 != int32(2619) {
						v312 = v308
					} else {
						v312 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
					m.G0 = v32 + int32(112)
					return
				} else {
					v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[9])))
					if v174 != int32(1) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
						v308 = int32(0)
						if l0 != int32(2619) {
							v312 = v308
						} else {
							v312 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
						m.G0 = v32 + int32(112)
						return
					} else {
						v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[10])))
						if v178&int32(1) == int32(0) {
							*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
							v308 = int32(0)
							if l0 != int32(2619) {
								v312 = v308
							} else {
								v312 = int32(0)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
							m.G0 = v32 + int32(112)
							return
						} else {
							v184 = float32(0)
							v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
							if base.F32_lt(v185, v184) != 0 {
								v188 = v184
							} else {
								v188 = v185
							}
							v189 = float32(1)
							v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
							if v190 <= int32(0) {
								v204 = v189
							} else {
								v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
								if v193 <= int32(0) {
									v204 = v189
								} else {
									if base.Ui32(v193) < base.Ui32(v190) {
										v198 = v193
									} else {
										v198 = v190
									}
									v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
								}
							}
							v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
							v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
							v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
							v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v111))
							if v110 < int32(0) {
								v219 = v212
							} else {
								v215 = base.F32_convert_i32_u(v110)
								if base.F32_gt(v212, v215) == int32(0) {
									v219 = v212
								} else {
									v219 = v215
								}
							}
							v221 = base.F32_convert_i64_s(v205)
							v222 = base.F32_convert_i64_s(v206)
							v223 = base.F32_convert_i64_s(v207)
							v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v114), v188), base.F32_convert_i32_s(v113))
							v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v115)), v204), base.F32_convert_i32_s(v109))
							v233 = F_errstart(m, int32(12), int32(0))
							mBase = m.M
							v234 = m.ExcPending
							if v234 != 0 {
								return
							} else {
								if int32(0) <= v109 {
									if v233 == int32(0) {
										if v170|base.F32_lt(v219, v223) != 0 {
											v289 = int32(1)
										} else {
											v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
										v308 = base.F32_gt(v221, v226)
										if l0 != int32(2619) {
											v312 = v308
										} else {
											v312 = int32(0)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
										m.G0 = v32 + int32(112)
										return
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
										*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
										F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_0), v32)
										mBase = m.M
										v257 = m.ExcPending
										if v257 != 0 {
											return
										} else {
											v278 = int32(3138)
											F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return
											} else {
												if v170|base.F32_lt(v219, v223) != 0 {
													v289 = int32(1)
												} else {
													v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
												v308 = base.F32_gt(v221, v226)
												if l0 != int32(2619) {
													v312 = v308
												} else {
													v312 = int32(0)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
												m.G0 = v32 + int32(112)
												return
											}
										}
									}
								} else {
									if v233 == int32(0) {
										if v170|base.F32_lt(v219, v223) != 0 {
											v289 = int32(1)
										} else {
											v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
										v308 = base.F32_gt(v221, v226)
										if l0 != int32(2619) {
											v312 = v308
										} else {
											v312 = int32(0)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
										m.G0 = v32 + int32(112)
										return
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
										*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
										*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
										F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_3), v32-int32(-64))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
											return
										} else {
											v278 = int32(3142)
											F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
											mBase = m.M
											v281 = m.ExcPending
											if v281 != 0 {
												return
											} else {
												if v170|base.F32_lt(v219, v223) != 0 {
													v289 = int32(1)
												} else {
													v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
												v308 = base.F32_gt(v221, v226)
												if l0 != int32(2619) {
													v312 = v308
												} else {
													v312 = int32(0)
												}
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
												m.G0 = v32 + int32(112)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v166 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v166)
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v166)
				m.G0 = v32 + int32(112)
				return
			}
		} else {
			v143 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v143)
			v170 = v143
			if l3 == int32(0) {
				*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
				v308 = int32(0)
				if l0 != int32(2619) {
					v312 = v308
				} else {
					v312 = int32(0)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
				m.G0 = v32 + int32(112)
				return
			} else {
				v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[9])))
				if v174 != int32(1) {
					*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
					v308 = int32(0)
					if l0 != int32(2619) {
						v312 = v308
					} else {
						v312 = int32(0)
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
					m.G0 = v32 + int32(112)
					return
				} else {
					v178 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_needs_vacanalyze[10])))
					if v178&int32(1) == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v170)
						v308 = int32(0)
						if l0 != int32(2619) {
							v312 = v308
						} else {
							v312 = int32(0)
						}
						*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
						m.G0 = v32 + int32(112)
						return
					} else {
						v184 = float32(0)
						v185 = *(*float32)(unsafe.Add(mBase, uint32(l2)+100))
						if base.F32_lt(v185, v184) != 0 {
							v188 = v184
						} else {
							v188 = v185
						}
						v189 = float32(1)
						v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
						if v190 <= int32(0) {
							v204 = v189
						} else {
							v193 = *(*int32)(unsafe.Add(mBase, uint32(l2)+108))
							if v193 <= int32(0) {
								v204 = v189
							} else {
								if base.Ui32(v193) < base.Ui32(v190) {
									v198 = v193
								} else {
									v198 = v190
								}
								v204 = base.F32_sub(float32(1), base.F32_div(base.F32_convert_i32_u(v198), base.F32_convert_i32_u(v190)))
							}
						}
						v205 = *(*int64)(unsafe.Add(mBase, uint32(l3)+88))
						v206 = *(*int64)(unsafe.Add(mBase, uint32(l3)+96))
						v207 = *(*int64)(unsafe.Add(mBase, uint32(l3)+80))
						v212 = base.F32_add(base.F32_mul(base.F32_demote_f64(v116), v188), base.F32_convert_i32_s(v111))
						if v110 < int32(0) {
							v219 = v212
						} else {
							v215 = base.F32_convert_i32_u(v110)
							if base.F32_gt(v212, v215) == int32(0) {
								v219 = v212
							} else {
								v219 = v215
							}
						}
						v221 = base.F32_convert_i64_s(v205)
						v222 = base.F32_convert_i64_s(v206)
						v223 = base.F32_convert_i64_s(v207)
						v226 = base.F32_add(base.F32_mul(base.F32_demote_f64(v114), v188), base.F32_convert_i32_s(v113))
						v230 = base.F32_add(base.F32_mul(base.F32_mul(v188, base.F32_demote_f64(v115)), v204), base.F32_convert_i32_s(v109))
						v233 = F_errstart(m, int32(12), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return
						} else {
							if int32(0) <= v109 {
								if v233 == int32(0) {
									if v170|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v308 = base.F32_gt(v221, v226)
									if l0 != int32(2619) {
										v312 = v308
									} else {
										v312 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+48)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+40)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+32)) = base.F64_promote_f32(v230)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+24)) = base.F64_promote_f32(v222)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+16)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32))) = l2 + int32(4)
									F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_0), v32)
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return
									} else {
										v278 = int32(3138)
										F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v170|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v308 = base.F32_gt(v221, v226)
											if l0 != int32(2619) {
												v312 = v308
											} else {
												v312 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
											m.G0 = v32 + int32(112)
											return
										}
									}
								}
							} else {
								if v233 == int32(0) {
									if v170|base.F32_lt(v219, v223) != 0 {
										v289 = int32(1)
									} else {
										v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
									v308 = base.F32_gt(v221, v226)
									if l0 != int32(2619) {
										v312 = v308
									} else {
										v312 = int32(0)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
									m.G0 = v32 + int32(112)
									return
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v32)+96)) = base.F64_promote_f32(v226)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+88)) = base.F64_promote_f32(v221)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+80)) = base.F64_promote_f32(v219)
									*(*float64)(unsafe.Add(mBase, uint32(v32)+72)) = base.F64_promote_f32(v223)
									*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = l2 + int32(4)
									F_errmsg_internal(m, int32(_a_F_relation_needs_vacanalyze_3), v32-int32(-64))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return
									} else {
										v278 = int32(3142)
										F_errfinish(m, int32(_a_F_relation_needs_vacanalyze_1), v278, int32(_a_F_relation_needs_vacanalyze_2))
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return
										} else {
											if v170|base.F32_lt(v219, v223) != 0 {
												v289 = int32(1)
											} else {
												v289 = base.B2i32(int32(0) <= v109) & base.F32_lt(v230, v222)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v289)
											v308 = base.F32_gt(v221, v226)
											if l0 != int32(2619) {
												v312 = v308
											} else {
												v312 = int32(0)
											}
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v312)
											m.G0 = v32 + int32(112)
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
}
func F_relation_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_LockRelationOid(m, l0, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_RelationIdGetRelation(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				if v12 != 0 {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
					if v15 == int32(116) {
						v18 = int32(_a_F_relation_open_0)
						v20 = *(*int32)(unsafe.Add(mBase, _c_F_relation_open[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_relation_open[0])) = v20 | int32(1)
					} else {
					}
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
					switch v26 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_open[1])))
						if v30 == int32(0) {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
							if v33 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = int32(0)
							} else {
							}
							v39 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
						} else {
							v36 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v36)
						}
					default:
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
					}
					m.G0 = v6 + int32(16)
					return v12
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
						F_errmsg_internal(m, int32(_a_F_relation_open_1), v6)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_relation_open_2), int32(61), int32(_a_F_relation_open_3))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
	} else {
		v12 = F_RelationIdGetRelation(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+118)))
				if v15 == int32(116) {
					v18 = int32(_a_F_relation_open_0)
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_relation_open[0]))
					*(*int32)(unsafe.Add(mBase, _c_F_relation_open[0])) = v20 | int32(1)
				} else {
				}
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
				switch v26 - int32(83) {
				case 0, 22, 26, 29, 31, 33:
					v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_open[1])))
					if v30 == int32(0) {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+272))
						if v33 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v33)+128)) = int32(0)
						} else {
						}
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
					} else {
						v36 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v36)
					}
				default:
					v39 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v39
					*(*uint8)(unsafe.Add(mBase, uint32(v12)+268)) = uint8(v39)
				}
				m.G0 = v6 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg_internal(m, int32(_a_F_relation_open_1), v6)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_relation_open_2), int32(61), int32(_a_F_relation_open_3))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
func F_relation_openrv(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int32(0)
			v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_RelationIdGetRelation(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 != 0 {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
						if v20 == int32(116) {
							v23 = int32(_a_F_relation_openrv_0)
							v25 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0])) = v25 | int32(1)
						} else {
						}
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
						switch v31 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv[1])))
							if v35 == int32(0) {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
								if v38 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = int32(0)
								} else {
								}
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
							} else {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v41)
							}
						default:
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
						}
						m.G0 = v6 + int32(16)
						return v17
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
							F_errmsg_internal(m, int32(_a_F_relation_openrv_1), v6)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_relation_openrv_2), int32(61), int32(_a_F_relation_openrv_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
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
	} else {
		v12 = int32(0)
		v15 = F_RangeVarGetRelidExtended(m, l0, l1, v12, v12, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = F_RelationIdGetRelation(m, v15)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+118)))
					if v20 == int32(116) {
						v23 = int32(_a_F_relation_openrv_0)
						v25 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv[0])) = v25 | int32(1)
					} else {
					}
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
					switch v31 - int32(83) {
					case 0, 22, 26, 29, 31, 33:
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv[1])))
						if v35 == int32(0) {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v17)+272))
							if v38 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = int32(0)
							} else {
							}
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
						} else {
							v41 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v41)
						}
					default:
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v44
						*(*uint8)(unsafe.Add(mBase, uint32(v17)+268)) = uint8(v44)
					}
					m.G0 = v6 + int32(16)
					return v17
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v15
						F_errmsg_internal(m, int32(_a_F_relation_openrv_1), v6)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_relation_openrv_2), int32(61), int32(_a_F_relation_openrv_3))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
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
func F_relation_openrv_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		F_ReceiveSharedInvalidMessages(m)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = int32(0)
			v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16 != 0 {
					v18 = F_RelationIdGetRelation(m, v16)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
								F_errmsg_internal(m, int32(_a_F_relation_openrv_extended_0), v8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_relation_openrv_extended_1), int32(61), int32(_a_F_relation_openrv_extended_2))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
							if v23 == int32(116) {
								v26 = int32(_a_F_relation_openrv_extended_3)
								v28 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0])) = v28 | int32(1)
							} else {
							}
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
							switch v34 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv_extended[1])))
								if v38 == int32(0) {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
									if v41 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = int32(0)
									} else {
									}
									v47 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
								} else {
									v44 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v44)
								}
							default:
								v47 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
							}
							v52 = v18
							m.G0 = v8 + int32(16)
							return v52
						}
					}
				} else {
					v52 = int32(0)
					m.G0 = v8 + int32(16)
					return v52
				}
			}
		}
	} else {
		v14 = int32(0)
		v16 = F_RangeVarGetRelidExtended(m, l0, l1, l2, v14, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 != 0 {
				v18 = F_RelationIdGetRelation(m, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v18 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v16
							F_errmsg_internal(m, int32(_a_F_relation_openrv_extended_0), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_relation_openrv_extended_1), int32(61), int32(_a_F_relation_openrv_extended_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
						if v23 == int32(116) {
							v26 = int32(_a_F_relation_openrv_extended_3)
							v28 = *(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_relation_openrv_extended[0])) = v28 | int32(1)
						} else {
						}
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+119)))
						switch v34 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_openrv_extended[1])))
							if v38 == int32(0) {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+272))
								if v41 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v41)+128)) = int32(0)
								} else {
								}
								v47 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
							} else {
								v44 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v44)
							}
						default:
							v47 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v18)+272)) = v47
							*(*uint8)(unsafe.Add(mBase, uint32(v18)+268)) = uint8(v47)
						}
						v52 = v18
						m.G0 = v8 + int32(16)
						return v52
					}
				}
			} else {
				v52 = int32(0)
				m.G0 = v8 + int32(16)
				return v52
			}
		}
	}
}
