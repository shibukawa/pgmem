package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_fetch_row_version(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+60)) = uint16(v10)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v12
	v15 = l3 + int32(52)
	v19 = F_heap_fetch(m, l0, l2, v15, v8+int32(12), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			F_ExecStorePinnedBufferHeapTuple(m, v15, l3, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v26
				m.G0 = v8 + int32(16)
				return v19
			}
		} else {
			m.G0 = v8 + int32(16)
			return v19
		}
	}
}
func F_heapam_relation_nontransactional_truncate(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_RelationTruncate(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_heapam_scan_sample_next_block(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v11 != 0 {
			F_ReleaseBuffer(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v20 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l1, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v46 = v20
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
						if v46 == int32(-1) {
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
							return v50
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
							if v55 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
										v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v65&int32(1) != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v65&int32(1) != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v22 == int32(-1) {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v46 = v25
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
						if v46 == int32(-1) {
							v50 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
							return v50
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
							if v55 != 0 {
								F_ProcessInterrupts(m)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
										v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v65&int32(1) != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									}
								}
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v65&int32(1) != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								}
							}
						}
					} else {
						v27 = v22 + int32(1)
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if base.Ui32(v27) < base.Ui32(v29) {
							v31 = v27
						} else {
							v31 = int32(0)
						}
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
						if v32&int32(128) != 0 {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							F_ss_report_location(m, v35, v31)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
								if v31 != v38 {
									v46 = v31
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
									if v46 == int32(-1) {
										v50 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
										return v50
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
										if v55 != 0 {
											F_ProcessInterrupts(m)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v59 = int32(0)
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
												v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
													v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
													if v65&int32(1) != 0 {
														F_heap_prepare_pagescan(m, l0)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return int32(0)
														} else {
															v70 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
															return v70
														}
													} else {
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
														return v70
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
												v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v65&int32(1) != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return int32(0)
													} else {
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
														return v70
													}
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
									v42 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v42)
									return v42
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v31 != v38 {
								v46 = v31
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
								if v46 == int32(-1) {
									v50 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
									return v50
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
												v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v65&int32(1) != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return int32(0)
													} else {
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
														return v70
													}
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
											v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v65&int32(1) != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v42)
								return v42
							}
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v20 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l1, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v46 = v20
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
					if v46 == int32(-1) {
						v50 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
						return v50
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
						if v55 != 0 {
							F_ProcessInterrupts(m)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v65&int32(1) != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v59 = int32(0)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
								if v65&int32(1) != 0 {
									F_heap_prepare_pagescan(m, l0)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								} else {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
									return v70
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v22 == int32(-1) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v46 = v25
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
					if v46 == int32(-1) {
						v50 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
						return v50
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
						if v55 != 0 {
							F_ProcessInterrupts(m)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v59 = int32(0)
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
								v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
									if v65&int32(1) != 0 {
										F_heap_prepare_pagescan(m, l0)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								}
							}
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v59 = int32(0)
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
								if v65&int32(1) != 0 {
									F_heap_prepare_pagescan(m, l0)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
										return v70
									}
								} else {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
									return v70
								}
							}
						}
					}
				} else {
					v27 = v22 + int32(1)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if base.Ui32(v27) < base.Ui32(v29) {
						v31 = v27
					} else {
						v31 = int32(0)
					}
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
					if v32&int32(128) != 0 {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_ss_report_location(m, v35, v31)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							if v31 != v38 {
								v46 = v31
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
								if v46 == int32(-1) {
									v50 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
									return v50
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
									if v55 != 0 {
										F_ProcessInterrupts(m)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v59 = int32(0)
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
											v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
												v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
												if v65&int32(1) != 0 {
													F_heap_prepare_pagescan(m, l0)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return int32(0)
													} else {
														v70 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
														return v70
													}
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
											v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v65&int32(1) != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
								v42 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v42)
								return v42
							}
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						if v31 != v38 {
							v46 = v31
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v46
							if v46 == int32(-1) {
								v50 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v50)
								return v50
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_sample_next_block[0]))
								if v55 != 0 {
									F_ProcessInterrupts(m)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v59 = int32(0)
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
										v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
											v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
											if v65&int32(1) != 0 {
												F_heap_prepare_pagescan(m, l0)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return int32(0)
												} else {
													v70 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
													return v70
												}
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										}
									}
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = int32(0)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
									v62 = F_ReadBufferExtended(m, v58, v59, v46, v59, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v62
										v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
										if v65&int32(1) != 0 {
											F_heap_prepare_pagescan(m, l0)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v70 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
												return v70
											}
										} else {
											v70 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v70)
											return v70
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(-1)
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v42)
							return v42
						}
					}
				}
			}
		}
	}
}
func F_heapam_tuple_complete_speculative(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v14)
	v17 = l1 + int32(32)
	v21 = F_ExecFetchSlotHeapTuple(m, l1, v14, v12+int32(15))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if l3 != 0 {
			v23 = m.G0
			v24 = int32(16)
			v25 = v23 - v24
			m.G0 = v25
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+2)))
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17))))
			v32 = F_ReadBuffer(m, l0, v27|v28<<(uint(v24)%32))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_LockBufferInternal(m, v32, int32(3))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v32 < int32(0) {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[0]))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v32^int32(-1))<<(uint(int32(2))%32))))
						v54 = v46
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[1]))
						v54 = v48 + v32<<(uint(int32(13))%32) + int32(-8192)
					}
					v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
					if v55 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_heapam_tuple_complete_speculative_0), int32(0))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_heapam_tuple_complete_speculative_1), int32(_a_F_heapam_tuple_complete_speculative_2), int32(_a_F_heapam_tuple_complete_speculative_3))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
						if base.Ui32(int32(25)) <= base.Ui32(v58) {
							v66 = int32(base.Ui32(v58+int32(_a_F_heapam_tuple_complete_speculative_4)) >> (uint(int32(2)) % 32))
						} else {
							v66 = int32(0)
						}
						if base.Ui32(v66&int32(_a_F_heapam_tuple_complete_speculative_5)) < base.Ui32(v55) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_heapam_tuple_complete_speculative_0), int32(0))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_heapam_tuple_complete_speculative_1), int32(_a_F_heapam_tuple_complete_speculative_2), int32(_a_F_heapam_tuple_complete_speculative_3))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55<<(uint(int32(2))%32))+20))
							if v73&int32(_a_F_heapam_tuple_complete_speculative_6) != int32(_a_F_heapam_tuple_complete_speculative_7) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v157 = m.ExcPending
								if v157 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_heapam_tuple_complete_speculative_8), int32(0))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_heapam_tuple_complete_speculative_1), int32(_a_F_heapam_tuple_complete_speculative_9), int32(_a_F_heapam_tuple_complete_speculative_3))
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v78 = int32(_a_F_heapam_tuple_complete_speculative_10)
								v80 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
								*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v80 + int32(1)
								F_MarkBufferDirty(m, v32)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									v88 = v54 + v73&int32(_a_F_heapam_tuple_complete_speculative_11)
									v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v88)+16)) = uint16(v89)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v91
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+118)))
									if v94 != int32(112) {
										v129 = int32(_a_F_heapam_tuple_complete_speculative_10)
										v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v131 - int32(1)
										F_UnlockReleaseBuffer(m, v32)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											m.G0 = v25 + int32(16)
											v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
											if v174 == int32(1) {
												F_pfree(m, v21)
												mBase = m.M
												v178 = m.ExcPending
												if v178 != 0 {
													return
												} else {
													m.G0 = v12 + int32(16)
													return
												}
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[3]))
										if v98 <= int32(0) {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
											if v101 != 0 {
												v129 = int32(_a_F_heapam_tuple_complete_speculative_10)
												v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v131 - int32(1)
												F_UnlockReleaseBuffer(m, v32)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													m.G0 = v25 + int32(16)
													v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
													if v174 == int32(1) {
														F_pfree(m, v21)
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													} else {
														m.G0 = v12 + int32(16)
														return
													}
												}
											} else {
												v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												if v102 != 0 {
													v129 = int32(_a_F_heapam_tuple_complete_speculative_10)
													v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v131 - int32(1)
													F_UnlockReleaseBuffer(m, v32)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return
													} else {
														m.G0 = v25 + int32(16)
														v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
														if v174 == int32(1) {
															F_pfree(m, v21)
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return
															} else {
																m.G0 = v12 + int32(16)
																return
															}
														} else {
															m.G0 = v12 + int32(16)
															return
														}
													}
												} else {
													v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
													*(*uint16)(unsafe.Add(mBase, uint32(v25)+14)) = uint16(v103)
													F_XLogBeginInsert(m)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return
													} else {
														v108 = int32(_a_F_heapam_tuple_complete_speculative_12)
														v110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[4])))
														v111 = v110 | int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[4])) = uint8(v111)
														F_XLogRegisterData(m, v25+int32(14), int32(2))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return
														} else {
															F_XLogRegisterBuffer(m, int32(0), v32, int32(8))
															mBase = m.M
															v121 = m.ExcPending
															if v121 != 0 {
																return
															} else {
																v124 = F_XLogInsert(m, int32(10), int32(80))
																mBase = m.M
																v125 = m.ExcPending
																if v125 != 0 {
																	return
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v54))) = base.I64_rotl(v124, int64(32))
																	v129 = int32(_a_F_heapam_tuple_complete_speculative_10)
																	v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
																	*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v131 - int32(1)
																	F_UnlockReleaseBuffer(m, v32)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return
																	} else {
																		m.G0 = v25 + int32(16)
																		v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
																		if v174 == int32(1) {
																			F_pfree(m, v21)
																			mBase = m.M
																			v178 = m.ExcPending
																			if v178 != 0 {
																				return
																			} else {
																				m.G0 = v12 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v12 + int32(16)
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
											v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
											*(*uint16)(unsafe.Add(mBase, uint32(v25)+14)) = uint16(v103)
											F_XLogBeginInsert(m)
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												v108 = int32(_a_F_heapam_tuple_complete_speculative_12)
												v110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[4])))
												v111 = v110 | int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[4])) = uint8(v111)
												F_XLogRegisterData(m, v25+int32(14), int32(2))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return
												} else {
													F_XLogRegisterBuffer(m, int32(0), v32, int32(8))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return
													} else {
														v124 = F_XLogInsert(m, int32(10), int32(80))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return
														} else {
															*(*int64)(unsafe.Add(mBase, uint32(v54))) = base.I64_rotl(v124, int64(32))
															v129 = int32(_a_F_heapam_tuple_complete_speculative_10)
															v131 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2]))
															*(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_complete_speculative[2])) = v131 - int32(1)
															F_UnlockReleaseBuffer(m, v32)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return
															} else {
																m.G0 = v25 + int32(16)
																v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
																if v174 == int32(1) {
																	F_pfree(m, v21)
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return
																	} else {
																		m.G0 = v12 + int32(16)
																		return
																	}
																} else {
																	m.G0 = v12 + int32(16)
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
					}
				}
			}
		} else {
			F_heap_abort_speculative(m, l0, v17)
			mBase = m.M
			v168 = m.ExcPending
			if v168 != 0 {
				return
			} else {
				v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
				if v174 == int32(1) {
					F_pfree(m, v21)
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return
					} else {
						m.G0 = v12 + int32(16)
						return
					}
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F_heapam_tuple_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)) = uint8(v18)
	v23 = F_ExecFetchSlotHeapTuple(m, l2, v18, v16+int32(15))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v27
		v30 = F_heap_update(m, l0, l1, v23, l3, l4, l6, l7, l8, l9, l10)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l2)+36)) = uint16(v32)
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v34
			if v30 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l10))) = int32(0)
			} else {
			}
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+15)))
			if v38 == int32(1) {
				F_pfree(m, v23)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					m.G0 = v16 + int32(16)
					return v30
				}
			} else {
				m.G0 = v16 + int32(16)
				return v30
			}
		}
	}
}
